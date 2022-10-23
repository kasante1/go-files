package main

import (
	"fmt"
)


func main(){

		willclose :=  make(chan int, 10)

		willclose <- -1
		willclose <- 0
		willclose <- 2

		<-willclose
		<-willclose
		<-willclose
		close(willclose)
		read := <-willclose
		fmt.Println(read)

	}

