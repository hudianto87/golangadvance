package main

import (
	"fmt"
	"time"
)

func givemeresponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "hello world"
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go givemeresponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
