package main

import (
	"belajargolangpart2/session4crud/entity"
	"belajargolangpart2/session4crud/handler"
	"belajargolangpart2/session4crud/repository/slice"
	"belajargolangpart2/session4crud/router"
	"belajargolangpart2/session4crud/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHanlder := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHanlder)

	log.Println("runnning service on port 8080")
	r.Run("localhost:8080")
}
