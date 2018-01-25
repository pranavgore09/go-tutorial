package main

import "fmt"

type Notifier interface {
	Notify()
}

type Email struct {
	id string
}

type SMS struct {
	number string
}

type Push struct{}

func (e Email) Notify() {
	fmt.Println("Hello, I will send email to ", e.id)
}

func (s SMS) Notify() {
	fmt.Println("Hello, I will send message to ", s.number)
}

func SendNotifications(n Notifier) {
	n.Notify()
}

func main() {
	e := Email{"pranav@rhd"}
	s := SMS{"+91987654321"}
	// p := Push{}
	SendNotifications(e)
	SendNotifications(s)
	// SendNotifications(p)
}
