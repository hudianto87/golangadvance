package main

import (
	"fmt"
	"time"
)

func onlyin(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "hello world"
}

func onlyout(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go onlyin(channel)
	go onlyout(channel)

	time.Sleep(5 * time.Second)
}
