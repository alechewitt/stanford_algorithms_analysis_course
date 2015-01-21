package main 

import (
	"fmt"
)

func mergeSort(mySlice[]int) []int{
	if len(mySlice) <= 1 {
		return mySlice
	} else {
		// mySliceResult := make([]int, 0, 1000)
		mySliceA := mySlice[:3]
		// mySliceB := [3:6]
		return mySliceA
	}	
}

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(mergeSort(p))
}