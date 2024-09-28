package main

import (
	"fmt"
	"sync"
	"time"
)

func RunAsync(group *sync.WaitGroup, a int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("hello", a)
	time.Sleep(1 * time.Second)
}

func main() {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group, i)
	}

	group.Wait()
	fmt.Println("Done")
}
