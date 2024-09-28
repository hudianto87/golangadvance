package router

import (
	"belajargolangpart2/session6dbpgx-crud/handler"
	"belajargolangpart2/session6dbpgx-crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler) {
	userPublichEndpoint := r.Group("/users")

	userPublichEndpoint.GET("/:id", userHandler.GetUser)
	userPublichEndpoint.GET("", userHandler.GetAllUsers)
	userPublichEndpoint.GET("/", userHandler.GetAllUsers)

	userPrivateEndpoint := r.Group("/users")
	userPrivateEndpoint.Use(middleware.AuthMiddleware())
	userPrivateEndpoint.POST("", userHandler.CreateUser)
	userPrivateEndpoint.POST("/", userHandler.CreateUser)
	userPrivateEndpoint.PUT("/:id", userHandler.UpdateUser)
	userPrivateEndpoint.DELETE("/:id", userHandler.DeleteUser)
}
