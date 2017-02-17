package main

import (
	"fmt"
	"time"
)

// Now we added 'GO' keyword before function call
// It will make that function run in a goroutine
// Try to run and see output. (Must run before going ahead)
func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo ", i)
		time.Sleep(4 * time.Millisecond)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar ", i)
		time.Sleep(5 * time.Millisecond)
	}
}
