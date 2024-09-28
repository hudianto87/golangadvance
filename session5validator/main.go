package main

import (
	"belajargolangpart2/session5validator/entity"
	"belajargolangpart2/session5validator/handler"
	"belajargolangpart2/session5validator/repository/slice"
	"belajargolangpart2/session5validator/router"
	"belajargolangpart2/session5validator/service"
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
