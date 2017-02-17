package main

import (
	"fmt"
	"time"
)

// Lets write to a channel and read from the same
// Why do we need to have a goroutine for writing to channel ?

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
