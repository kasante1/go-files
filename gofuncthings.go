package main
import (
	"fmt"
	"sort"
)


type Person struct{
	Firstname string
	Lastname string
	Age int
}


func main(){
	people := []Person{
		{"pat","patterson",37},
		{"tracy","bobbert",23},
		{"fred","fredson",18},
	}
	

	sort.Slice(people, func(i int, j int) bool{
		return people[i].Lastname < people[j].Lastname
	})
	fmt.Println(people)
}