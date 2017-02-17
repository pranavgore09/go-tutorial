package main

import (
	"fmt"
	"time"
)

// Try the Bomb

func main() {
	boom := time.After(5 * time.Second)
	for {
		select {
		case <-boom:
			fmt.Println("* !! BOOM !! *")
			return
		default:
			fmt.Println("tick .")
			time.Sleep(1 * time.Second)
		}
	}
}

// Do same thing with Ticker.C !
// func main() {
// 	boom := time.After(5 * time.Second)
// 	ticker := time.NewTicker(1 * time.Second)
// 	for {
// 		select {
// 		case <-boom:
// 			fmt.Println("* !! BOOM !! *")
// 			return
// 		case <-ticker.C:
// 			fmt.Println("tick .")
// 		}
// 	}
// }
