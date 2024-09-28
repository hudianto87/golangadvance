package router

import (
	"belajargolangpart2/session1/handler"
	"belajargolangpart2/session1/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	//yang ini bisa dipanggil tanpa authorization
	r.GET("/", handler.RootHandler)

	//membuat grouping handler
	privateRoute := r.Group("/api/v1")
	privateRoute.Use(middleware.AuthMiddleware())
	{
		privateRoute.POST("/post", handler.PostHandler)
	}
}
