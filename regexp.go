package main

import (
	"fmt"
	s "strings"
)

func main() {
	var f = fmt.Printf

	upper := s.ToUpper("hello world")

	f("to upper : %s\n", upper)
	f("EqualFold : %v\n", s.EqualFold("mihalis","Mihalis"))
	f("are they equal : %v\n",s.EqualFold("mihalis","mihali"))
}