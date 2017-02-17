package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(err string, val int) {
	time.Sleep(3 * time.Millisecond)
	if err != "" {
		errorChan <- err
	} else {
		data <- val
	}
	wg.Done()
}

var data chan int
var errorChan chan string
var wg sync.WaitGroup

func main() {
	data = make(chan int)
	errorChan = make(chan string)
	wg.Add(3)
	go worker("Not found", 100)
	go worker("", 200)
	go worker("", 300)

	go func() {
		wg.Wait()
		close(data)
	}()
	for {
		select {
		case e := <-errorChan:
			fmt.Println("Error : ", e)
			return
		case x, ok := <-data:
			if ok == false {
				fmt.Println("Process completed without any error")
				return
			}
			fmt.Println(x)
		}
	}
}
