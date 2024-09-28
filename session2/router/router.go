package router

import (
	"belajargolangpart2/session2/handler"
	"belajargolangpart2/session2/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/", handler.GetAllUser)
	usersPublicEndpoint.GET("/:id", handler.GetUser)

	usersPrivateEndpoint := r.Group("/users")
	usersPrivateEndpoint.Use(middleware.AuthMiddleware()) // You might want to add middleware here
	usersPrivateEndpoint.POST("/", handler.CreatedUser)
	usersPrivateEndpoint.PUT("/:id", handler.UpdateUser)
	usersPrivateEndpoint.DELETE("/:id", handler.DeleteUser)
}
