package main 

import (
	"fmt"
)

func mergeSort(mySlice[]int) []int{
	if len(mySlice) <= 1 {
		return mySlice
	} else {
		// mySliceResult := make([]int, 0, 1000)
		halfLen := len(mySlice) / 2
		mySliceA := mergeSort(mySlice[:halfLen])
		mySliceB := mergeSort(mySlice[halfLen:])

		mySliceResult := make([]int, 0, len(mySlice))
		for len(mySliceA) > 0 || len(mySliceB) > 0 {
			if len(mySliceA) == 0 {
				mySliceResult = append(mySliceResult, mySliceB[0])
				mySliceB = mySliceB[1:]
			} else if len(mySliceB) == 0 {
				mySliceResult = append(mySliceResult, mySliceA[0])
				mySliceA = mySliceA[1:]
			} else {
				if mySliceA[0] > mySliceB[0] {
					mySliceResult = append(mySliceResult, mySliceB[0])
					mySliceB = mySliceB[1:]
				} else {
					mySliceResult = append(mySliceResult, mySliceA[0])
					mySliceA = mySliceA[1:]
				}
			}

		}
		return mySliceResult
	}	
}

func main() {
	p := []int{300, 200, 1, 2, 3, 5, 7, 11, 13, 1, 2}
	fmt.Println(mergeSort(p))
}