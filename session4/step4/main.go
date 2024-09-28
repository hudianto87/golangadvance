package main

import (
	"belajargolangpart2/session4/step4/entity"
	"belajargolangpart2/session4/step4/handler"
	"belajargolangpart2/session4/step4/repository/slice"
	"belajargolangpart2/session4/step4/router"
	"belajargolangpart2/session4/step4/service"
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
