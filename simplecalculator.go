package main

import (
	"fmt"
	"strconv"
)

type doubleFunc func(int,int)int

func add(i,j int)int {return i+j}
func sub(i,j int)int {return i-j}
func mul(i,j int)int {return i*j}
func div(i,j int)int {return i/j}

var opMap = map[string] doubleFunc{
	"+":add,
	"-":sub,
	"*":mul,
	"/":div,
}
func main(){
	expressions := [][]string{
		{"2","+","3"},
		{"2","-","3"},
		{"2","*","3"},
		{"2","/","3"},
		{"2","%","3"},
		{"two","+","three"},
		{"5"},
	}

	for _, expression := range expressions{
		if len(expression) != 3{
			fmt.Println("invalid expression: ", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil{
			fmt.Println(err)
			continue
		}
		
		op:= expression[1]
		opFunc, ok := opMap[op]
		if !ok{
			fmt.Println("unsupported operator: ",op)
			continue
		}
	
		p2,err := strconv.Atoi(expression[2])
		if err != nil{
			fmt.Println(err)
			continue
		}
		result := opFunc(p1,p2)
		fmt.Println(result)

	}
}