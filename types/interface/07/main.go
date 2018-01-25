package main

import "fmt"

type Notifier interface {
	Notify()
}

type Email struct {
	id string
}

type User struct {
	Email
	name string
}

func SendNotifications(n []Notifier) {
	for _, x := range n {
		x.Notify()
	}
}

func (e Email) Notify() {
	fmt.Println("Hello, I will send email to ", e.id)
}

func (u User) Notify() {
	fmt.Println("Hello, I will notify ", u.name)
}

func main() {
	e := Email{"pranav@rhd"}
	u := User{e, "Pranav Gore"}
	SendNotifications([]Notifier{e, u})
}
