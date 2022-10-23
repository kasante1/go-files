package main
import(
	"fmt"
	"sort"
)

func main(){
	input := []string{"foo","bar","baz"}

	var result []string
	result = append(input, "abc")
	result = append(result,"def")
	sort.Sort(sort.StringSlice(result))
	fmt.Print(result)
}