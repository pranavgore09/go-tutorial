package main

// interfaces are a way to define behaviour - using methods
// http://stackoverflow.com/questions/27775376/value-receiver-vs-pointer-receiver-in-golang

import (
	"fmt"
)

type runner interface {
	run()
}

type crawler interface {
	crawl()
}

type animal struct {
	intelegence int
	speed       int
}

type reptile struct {
	speed  int
	length int
}

func (r reptile) crawl() {
	fmt.Printf("I do Crawl at %d km/h\n", r.speed)
}

func (a animal) run() {
	fmt.Printf("I do run at %d km/h\n", a.speed)
}

func main() {
	fmt.Println("Understand go interfaces")
	cheetah := animal{10, 10}
	tiger := animal{10, 9}
	lion := animal{10, 8}
	blackMamba := reptile{10, 6}
	ghonas := reptile{20, 4}
	runners := []runner{cheetah, tiger, lion}
	for _, x := range runners {
		x.run()
	}
	crawlers := []crawler{blackMamba, ghonas}
	for _, x := range crawlers {
		x.crawl()
	}
}
