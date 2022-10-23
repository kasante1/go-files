package main

import (
	"fmt"
	"collections/deque"
)

func remove (s [] string, i int)[]string {
	h :=copy(s[i:],s[i+1:])
	fmt.Println(h)
	return s[:len(s)-1]
}

func main(){
	slice := []string{"foo","bar","baz","bass","tri"}
	//slice = append(slice, "tri")
	//fmt.Println(slice[1])
	remmoveElement := remove(slice,2)

	fmt.Println(remmoveElement)
	//so in golang the slice[] returns the entire slice

}