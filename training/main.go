package main

import (
	"fmt"
)

func main() {
	mp := map[interface{}]interface{}{
		"string": "string",
		1:        10,
		"float":  1.1,
	}
	fmt.Println(mp)
}
