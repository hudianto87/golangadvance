package main

import (
	"belajargolangpart2/session2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.SetupRouter(r)

	r.Run("localhost:8080")
}
