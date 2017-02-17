package main

import (
	"fmt"
	"time"
)

// Let's add more goroutines those will write to same channel
// see what happens, learn how to make use of previous knowledge
// What wil happen when we do not know "count of writers" ?

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
	}()
	go func() {
		for i := 10; i < 20; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
	}()
	go func() {
		for i := 20; i < 30; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
	}()
	for i := 0; i < 30; i++ {
		fmt.Println(<-ch)
	}
}
