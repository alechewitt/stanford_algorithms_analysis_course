// This program takes input
package main

import (
	"fmt"
)

func findIthOrderStatistic(inputSlice[]int, orderStat int) int{ 
	if (len(inputSlice) == 1) {
		return inputSlice[0]
	} else {
		// Choose a pivot,
		// take first element as pivot for simplicity, 
		// however better to take randomly from choice of three
		pivotIndex := 0
		// Partition input slice around pivot, exactly the same as quick sort

		i := 1
		for j:=1; j<len(inputSlice); j++ {
			if inputSlice[j] < inputSlice[pivotIndex] {
				inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
				i += 1
			}
		}
		// The pivots right full position
		pivotCorrectPosition := i - 1
		// Swap pivot into its righful place
		inputSlice[pivotIndex], inputSlice[pivotCorrectPosition] = inputSlice[pivotCorrectPosition], inputSlice[pivotIndex]
		if pivotCorrectPosition == orderStat{
			// We have found the number at i th order statistic that we are looking for
			return inputSlice[pivotCorrectPosition]
		} else if (pivotCorrectPosition > orderStat) {
			return findIthOrderStatistic(inputSlice[:pivotCorrectPosition], orderStat)
		} else {
			// Pivot right position is less than the order statistic we are looking for
			return findIthOrderStatistic(inputSlice[pivotCorrectPosition + 1:], orderStat - pivotCorrectPosition - 1)
		}
	}
}

func main() {
	// testSlice := []int{5,6,7,8}
	// testSlice := []int{5,8,7,6}
	testSlice := []int{6,5,8,7}
	fmt.Println(findIthOrderStatistic(testSlice, 0))
}