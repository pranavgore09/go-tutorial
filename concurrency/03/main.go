package main

import (
	"fmt"
	"time"
)

// In previous program it does not print anything because our main goroutine exits before other goroutines are completed
// So in that case we were not able to see anything on the console
// Let's add some Sleep in main goroutine so that it will sleep for sometime and let's hope other goroutines
// will complete their printing work, so that we can see expected output on console

func main() {
	go foo()
	go bar()
	time.Sleep(3 * time.Second)
	// But yeah, even I did not like this approach. Let go ahead to better future.
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo ", i)
		time.Sleep(5 * time.Millisecond)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar ", i)
		time.Sleep(3 * time.Millisecond)
	}
}
