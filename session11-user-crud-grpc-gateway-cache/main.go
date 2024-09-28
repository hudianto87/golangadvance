package main

import (
	grpcHanlder "belajargolangpart2/session11-user-crud-grpc-gateway-cache/handler/grpc"
	"belajargolangpart2/session11-user-crud-grpc-gateway-cache/middleware"
	pb "belajargolangpart2/session11-user-crud-grpc-gateway-cache/proto/user_service/v1"
	postgresgormraw "belajargolangpart2/session11-user-crud-grpc-gateway-cache/repository/postgres_gorm_raw"
	"belajargolangpart2/session11-user-crud-grpc-gateway-cache/service"
	"context"
	"log"
	"net"

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

	dsn := "postgresql://postgres:P4ssw0rd@192.168.26.50:5432/traininggolang"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatalln(err)
	}
	pgxpool.New(context.Background(), dsn)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

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

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		log.Println("running server on port : 50051")
		grpcServer.Serve(lis)
	}()

	// run grpc gateway
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	ginServer.Run("localhost:8080")

}
