package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}


func (u user)notify(){
	fmt.Printf("sending user email to %s<%s>\n",u.name,u.email)
}

func (u *user) changeEmail(email string){
	u.email = email
}

func sendNotification(n notifier){
	n.notify()
}

func main(){
	bill := user{"bill","bill@gmail.com"}

	bill.notify()
	bill.changeEmail("billnew@gmail.com")

	bill.notify()

	sendNotification(bill)
}

 