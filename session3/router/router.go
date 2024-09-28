package router

import (
	"belajargolangpart2/session3/handler"
	"belajargolangpart2/session3/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	//yang ini bisa dipanggil tanpa authorization
	r.GET("/", handler.RootHandler)

	//membuat grouping handler
	privateEndpoint := r.Group("/private")
	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		privateEndpoint.POST("/post", handler.PostHandler)
	}
}
