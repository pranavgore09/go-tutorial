package main

import (
	"fmt"
)

// Lets see a simplest possible solution with channels.
// Declare a channel, write to it, read from it.
// Simple, right ? Execute and see the output

func main() {
	ch := make(chan int)
	ch <- 10
	fmt.Println(<-ch)
}

// yes, it shows a deadlock
