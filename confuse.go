package main

import(
	"fmt"
	//"sort"
)

func findMinInt(a []int) int{
	minInt := a[0]

	for _, i := range a {
		if minInt > i{
			minInt = i
		}
	}
	return minInt
}

func main(){
	intData := []int{3,1,2,5,6,8,77,54,35,3,61,6,56,5}
	// stringData := []string{"foo","bar","baz","bass"}
	// floatData := []float64{1.5,3.6,7.9,2.5,10.6}

	// sort.Ints(intData)
	// sort.Strings(stringData)
	// sort.Float64s(floatData)

	// fmt.Println("sorted integers: ", intData, "\nSorted Strings: ",stringData,"\nSorted Floats: ",floatData)

	minResult := findMinInt(intData)
    fmt.Println(minResult)

	
}