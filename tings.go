package main

import(

	"fmt"

)

type myStructure struct{
	Name string
	Surname string
	Height int32
}

func createStruct (n,s string, h int32) *myStructure{
	if h > 10 {
		h = 0
	}
	return &myStructure{n,s,h}
}


func main() {

	s1 := createStruct("hello","heya",9)
	s2 := createStruct("hi","howdy",9)

	fmt.Println((*s1).Name)
	fmt.Println((s2.Height))
	fmt.Println(&s1)
	
}