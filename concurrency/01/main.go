package main

import (
	"fmt"
	"time"
)

// Following main function calls two functions back to back
// This program is running sequential and hence prints all statements back to back
func main() {
	foo()
	bar()
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
