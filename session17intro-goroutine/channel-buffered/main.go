package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("menerima data ke ", data)
	}
	time.Sleep(2 * time.Second)
}
