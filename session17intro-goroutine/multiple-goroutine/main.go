package main

import (
	"fmt"
	"time"
)

func displaynumber(number int) {
	fmt.Println("Display no", number)
}

func main() {
	start := time.Now()

	for i := 0; i < 100000; i++ {
		go displaynumber(i)
	}

	duration := time.Since(start)
	fmt.Printf("exec time : %v\n", duration)
	time.Sleep(3 * time.Second)
}
