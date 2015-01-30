package main

import (
	"fmt"
)

func calcInversions(list []int) int{
	numInversions := 0;
	for i:=0; i<len(list); i++ {
		for j:=i+1; j<len(list); j++ {
			if list[i] > list[j] {
				numInversions += 1
			}
		}
	}
	return numInversions;
}

func main() {
	testList := []int{6, 5, 4, 3, 2, 1}
	fmt.Println(calcInversions(testList))
}