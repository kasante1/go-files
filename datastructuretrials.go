package main
import(
	"fmt"
	"os"

)
// functions that returns other functions

func varFunc(input ...string){
	fmt.Println(input)
}

func oneByone(messages string, s ...int)int{
	fmt.Println(messages)
	sum:= 0
	for i, a := range s{
		fmt.Println(i, a)
		sum = sum + a
	}
	return sum
}
func main(){
	arguments := os.Args
	varFunc(arguments...)
	sum := oneByone("adding numbers...",1,2,3,4,5)
	fmt.Println(sum)
}
