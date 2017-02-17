package main

import (
	"fmt"
	"sync"
	"time"
)

// We will try to avoid that random Sleep time in main and hoping that
// other goroutines will complete their task within that time

// Let's have a int value that will be incremented when every goroutine is started
// and decremented when every goroutine exits
// When that count drops to zero, main will understand that all goroutines are exited.
// Yes, correct guess, we need to use Mutex while updating that int value.

var goroutineCnt int
var goroutineCntLock sync.Mutex

func incrementCnt() {
	goroutineCntLock.Lock()
	goroutineCnt++
	goroutineCntLock.Unlock()
}

func decrementCnt() {
	goroutineCntLock.Lock()
	goroutineCnt--
	goroutineCntLock.Unlock()
}

func main() {
	incrementCnt()
	go foo()
	incrementCnt()
	go bar()
	for {
		if goroutineCnt == 0 {
			break
		}
	}
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo ", i)
		time.Sleep(3 * time.Millisecond)
	}
	decrementCnt()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar ", i)
		time.Sleep(4 * time.Millisecond)
	}
	decrementCnt()
}
