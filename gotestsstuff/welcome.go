package main

import (
	"fmt"
)

func Hello(name, language string)string{
	return fmt.Sprint(greeting(language), name,)
}

var greetings = map[string]string{
	"es":"Hola ",
	"fr":"Bonjour ",
}

func greeting(language string)string{
	greeting, ok := greetings[language]
	if ok{
		return greeting
	}
	return "Hello "
}

func Welcome()string{
	return "Hello, world!"
}

func main(){
	greetsomeone := Hello("john","fr",)
	fmt.Println(greetsomeone)
	fmt.Println(Welcome())

	x := 5
	y := 9

	Swap(&x, &y)

	fmt.Println(x, y)

}