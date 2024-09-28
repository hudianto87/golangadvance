package main

import (
	"belajargolangpart2/session3/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.SetupRouter(r)

	r.Run(":8080")
}
