package main

import (
	"belajargolangpart2/session6dbpgx-crud/handler"
	postgrespgx "belajargolangpart2/session6dbpgx-crud/repository/postgres_pgx"
	"belajargolangpart2/session6dbpgx-crud/router"
	"belajargolangpart2/session6dbpgx-crud/service"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// go get itu per project nya, kalau go install untuk global

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	dsn := "postgresql://postgres:P4ssw0rd@192.168.26.50:5432/traininggolang"
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	pgxpool.New(context.Background(), dsn)

	userRepo := postgrespgx.NewUserRepository(pool)
	userService := service.NewUserService(userRepo)
	userHanlder := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHanlder)

	log.Println("runnning service on port 8080")
	r.Run("localhost:8080")
}
