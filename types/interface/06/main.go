package main

import "fmt"

type Notifier interface {
	Notify()
}

type Email struct {
	id string
}

func (e *Email) Notify() {
	fmt.Println("Hello, I will send email to ", e.id)
}

func main() {
	e := Email{"pranav@rhd"}
	n := []Notifier{&e}
	for _, x := range n {
		x.Notify()
	}
}
