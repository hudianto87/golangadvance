package main

import (
	"context"
	"log"
	"net"
	"session16-crud-user-docker-compose/entity"
	grpcHanlder "session16-crud-user-docker-compose/handler/grpc"
	"session16-crud-user-docker-compose/middleware"
	pb "session16-crud-user-docker-compose/proto/user_service/v1"
	postgresgormraw "session16-crud-user-docker-compose/repository/postgres_gorm_raw"
	"session16-crud-user-docker-compose/service"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	dsn := "postgresql://postgres:password@pg-db:5432/go_db_crud"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatalln(err)
	}
	pgxpool.New(context.Background(), dsn)

	//setup redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-db:6379",
		Password: "redispass",
		DB:       0,
	})

	//run migration
	if err := gormDB.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalln("Failed to migrate database : ", err)
	}

	//main -> handler -> service -> repository

	//untuk menggunakan gorm
	//userRepo := postgresgorm.NewUserRepository(gormDB)

	//untuk menggunakan raw
	userRepo := postgresgormraw.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo, rdb)
	userHanlder := grpcHanlder.NewUserHandler(userService)

	//run grpc server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor()))
	pb.RegisterUserServiceServer(grpcServer, userHanlder)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		log.Println("running server on port : 50051")
		grpcServer.Serve(lis)
	}()

	// run grpc gateway
	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server :", err)
	}

	gwmux := runtime.NewServeMux()

	if err = pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gatway :", err)
	}

	// run gin server
	ginServer := gin.Default()

	ginServer.Group("/v1/*{grpc_gatway").Any("", gin.WrapH(gwmux))

	log.Println("Running grpc gateway server in port :8080")

	ginServer.Run("0.0.0.0:8080")

}
