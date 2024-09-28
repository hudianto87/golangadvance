package router

import (
	"belajargolangpart2/session4/step4/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler) {
	userPublichEndpoint := r.Group("/users")

	userPublichEndpoint.GET("/", userHandler.GetAllUsers)
}
