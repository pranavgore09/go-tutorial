package main

import "fmt"

type Email struct {
	id string
}

type SMS struct {
	number string
}

func (e Email) Notify() {
	fmt.Println("Hello, I will send email to ", e.id)
}

func (s SMS) Notify() {
	fmt.Println("Hello, I will send message to ", s.number)
}

func main() {
	e := Email{"pranav@rhd"}
	e.Notify()

	s := SMS{"+91987654321"}
	s.Notify()
}
