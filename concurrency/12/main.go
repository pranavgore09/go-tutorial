package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go func() {
		time.Sleep(4 * time.Millisecond)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(3 * time.Millisecond)
		ch2 <- true
	}()

	select {
	case x := <-ch1:
		fmt.Println(x)
	case x := <-ch2:
		fmt.Println(x)
	}
}
