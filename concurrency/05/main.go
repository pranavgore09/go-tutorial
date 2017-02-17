package main

import (
	"fmt"
	"sync"
	"time"
)

// Let's use better approach.
// Go provides us sync.WaitGroup to do that job

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go foo()
	wg.Add(1)
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo ", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar ", i)
		time.Sleep(4 * time.Millisecond)
	}
	wg.Done()
}
