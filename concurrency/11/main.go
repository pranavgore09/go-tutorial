package main

import (
	"fmt"
	"sync"
	"time"
)

// But yeah, you are right -> we have WaitGroups built into golang !
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 10; i < 20; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 20; i < 25; i++ {
			ch <- i
			time.Sleep(time.Duration(i) * time.Millisecond)
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}
}
