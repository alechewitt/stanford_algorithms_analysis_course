package main

import(
	"fmt"
)

func countInversions(aList []int) (int, []int) {
	if len(aList) == 1 {
		return 0, aList
	}

	halfList := len(aList) / 2
	a, listA := countInversions(aList[0:halfList])
	b, listB := countInversions(aList[halfList:])

	listC := make([]int, 0, len(aList))
	c := 0
	for len(listA) > 0 || len(listB) > 0 {
		if len(listA) > 0 && len(listB) > 0 {
			if listA[0] <= listB[0] {
				// No inversion
				listC = append(listC, listA[0])
				listA = listA[1:]
			} else {
				// Inversion 
				c += len(listA)
				listC = append(listC, listB[0])
				listB = listB[1:]
			}
		} else if len(listA) > 0 {
			// List B has no elements left but A is 
			// still full of elements. We have counted all inversions
			// but still need to append the rest of listA to listC
			listC = append(listC, listA...)
			listA = listA[:0]
		} else {
			listC = append(listC, listB...)
			listB = listB[:0]
		}
	}
	totalInversions := a + b + c
	return totalInversions, listC
}

func main() {
	testSlice := []int{7,2,3,4,5,6,1}
	numInversions, newList := countInversions(testSlice)
	fmt.Println("Original list =", testSlice)
	fmt.Println("Number Inversions =", numInversions)
	fmt.Println("New Slice = ", newList)
}