package main

import (
	"fmt"
	"runtime"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	fmt.Printf("Hello, World\n")

	myarray := [3]string{"Jack", "Anne", "Jerry"}

	myslice := myarray[1:3]

	for _, value := range myslice {
		fmt.Printf("%s\r\n", value)
	}

	u := &user{"Bill", "bill@email.com"}

	sendNotification(u)

	u2 := admin{"Lisa", "lisa@email.com"}
	sendNotification(u2)

	fmt.Printf("CPU: %d\n", runtime.NumCPU())
}
