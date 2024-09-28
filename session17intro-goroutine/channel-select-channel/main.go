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
	channel1 := make(chan string)
	channel2 := make(chan string)

	go givemeresponse(channel1)
	go givemeresponse(channel2)

	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("data dari channel1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel2 ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
