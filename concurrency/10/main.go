package main

import (
	"fmt"
	"time"
)

// Multiple goroutines writing to same channel
// Not sure how many workers are there, not sure how much data will be written ?
// Let's use another channel and sort these things out

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
		done <- true
	}()
	go func() {
		for i := 10; i < 20; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
		done <- true
	}()
	go func() {
		for i := 20; i < 100; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		<-done
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}
}

// yes there is better approach, let's see in next snippet
