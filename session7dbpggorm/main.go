package main

import (
	"belajargolangpart2/session7dbpggorm/handler"
	postgresgormraw "belajargolangpart2/session7dbpggorm/repository/postgres_gorm_raw"
	"belajargolangpart2/session7dbpggorm/router"
	"belajargolangpart2/session7dbpggorm/service"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	dsn := "postgresql://postgres:P4ssw0rd@192.168.26.50:5432/traininggolang"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatalln(err)
	}
	pgxpool.New(context.Background(), dsn)

	//untuk menggunakan gorm
	//userRepo := postgresgorm.NewUserRepository(gormDB)

	//untuk menggunakan rawa
	userRepo := postgresgormraw.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHanlder := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHanlder)

	log.Println("runnning service on port 8080")
	r.Run("localhost:8080")
}
