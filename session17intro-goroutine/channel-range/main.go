package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "hello"
		channel <- "world"
		channel <- "!"
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-channel)
	}

	time.Sleep(2 * time.Second)
}
