package main

import (
	"belajargolangpart2/session1/router"

	"github.com/gin-gonic/gin"
)

// go get itu per project nya, kalau go install untuk global
// dari main -> handler -> service -> repository

func main() {

	//inisialisasi mode release
	gin.SetMode(gin.ReleaseMode)

	//inisialisasi router gin
	r := gin.Default()

	router.SetupRouter(r)

	//menjalankan server pada port 8080
	r.Run(":8080")
}
