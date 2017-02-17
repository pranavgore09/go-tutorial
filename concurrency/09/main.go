package main

import (
	"fmt"
	"time"
)

// Let's see how RANGE over channel works

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
		close(ch)
	}()
	for x := range ch {
		fmt.Println(x)
	}
}
