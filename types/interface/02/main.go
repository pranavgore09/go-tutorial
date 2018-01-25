package main

import "fmt"

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

func SendNotifications(i interface{}, typeName string) {
	switch typeName {
	case "email":
		e := i.(Email)
		e.Notify()
	case "sms":
		s := i.(SMS)
		s.Notify()
	default:
		fmt.Println("This type is not supported.")
	}
}

func main() {
	e := Email{"pranav@rhd"}
	s := SMS{"+91987654321"}
	p := Push{}
	SendNotifications(e, "email")
	SendNotifications(s, "sms")
	SendNotifications(p, "push")
}
