package main

import (
	"fmt"
	"time"
)

func helloworld() {
	fmt.Println("Hello world")
}

func main() {
	go helloworld()
	fmt.Println("main function")

	time.Sleep(1 * time.Second)
}
