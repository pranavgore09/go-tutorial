package main

import (
	"fmt"
)

// Simplest example of Buffered channel.
// It has slots to fill. once full it will wait for writing.
// notice that following code does not have any goroutine.
// Because read & write is not blocked until buffer is full

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Println(<-ch)
	ch <- 30
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
