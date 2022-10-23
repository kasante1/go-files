package main

import "fmt"

func incrementCounter() func() int{
	var initializedNumber = 0
	return func() int{
		initializedNumber++
		return initializedNumber
	}
}

func main(){
	n1 := incrementCounter()
	fmt.Println("n1 increment counter #1: ", n1())
	fmt.Println("n1 increment counter#2: ", n1())
	n2 := incrementCounter()
	fmt.Println("n2 increment counter#1: ", n2())
	fmt.Println("n1 increment counter#3: ",n1())
	fmt.Println("n2 increment counter#2: ",n2())
}