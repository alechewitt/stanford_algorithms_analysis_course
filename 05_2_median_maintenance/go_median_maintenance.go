package main 

import (
	"container/heap"
	"fmt"
	// "math"
)

/**
 * General heap methods
 */
type IntHeap []int

func (h IntHeap) Len() int { 
	return len(h) 
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i] 
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Max Heap Implementation
 */
type MaxHeap struct {
	IntHeap
}

func (h MaxHeap) Less(i, j int) bool { 
	return h.IntHeap[i] > h.IntHeap[j] 
}

func (h MaxHeap) Max() int {
	return h.IntHeap[0]
}

/**
 * Min Heap Implementation
 */
type MinHeap struct {
	IntHeap
}

func (h MinHeap) Less(i, j int) bool {
	return h.IntHeap[i] < h.IntHeap[j] 
}

func (h MinHeap) Min() int {
	return h.IntHeap[0]
}

/**
 * Main method
 */
func main() {
	testInput := []int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
	// Init upper half of heap that support Min operation
	upperHalf := &MinHeap{[]int{}}
	heap.Init(upperHalf)

	// Initialise lower half heap which supports max operation
	lowerHalf := &MaxHeap{[]int{}}
	heap.Init(lowerHalf)
    
	sumOfMedians := 0
	medians := []int{}
    for k := 0; k < len(testInput); k++ {
    	if k == 0 {
    		// Heaps currently dont contain anythin
    		heap.Push(lowerHalf, testInput[0])
    		medians = append(medians, testInput[0])
    		sumOfMedians += testInput[0]
    	} else if k == 1 {
    		if lowerHalf.Max() > testInput[1] {
    			heap.Push(upperHalf, heap.Pop(lowerHalf))
    			heap.Push(lowerHalf, testInput[1])
    			sumOfMedians += testInput[1]
    			medians = append(medians, testInput[1])
    		} else {
    			// New number is larger or the same, therefore append to 
    			// upperHalf heap
    			heap.Push(upperHalf, testInput[1])
    			sumOfMedians += lowerHalf.Max()
    			medians = append(medians, lowerHalf.Max())
    		}
    	} else {
    		// Heaps both contain at least one element
    	    if testInput[k] > lowerHalf.Max() {
    	    	if lowerHalf.Len() >= upperHalf.Len() {
	    	    	heap.Push(upperHalf, testInput[k])
    	    	} else {
	    	    	// Push into upper halfHeap
    	    		heap.Push(lowerHalf, heap.Pop(upperHalf))
    	    		heap.Push(upperHalf, testInput[k])
    	    	}
    	    } else {
    	    	// testInput[k] <= lowerHalf.Max()
    	    	if lowerHalf.Len() <= upperHalf.Len() {
    	    		heap.Push(lowerHalf, testInput[k])
    	    	} else {
    	    		heap.Push(upperHalf, heap.Pop(lowerHalf))
    	    		heap.Push(lowerHalf, testInput[k])
    	    	}
    	    }

    	    // Now determine which number is median
    	    lowerLen := lowerHalf.Len()
    	    upperLen := upperHalf.Len()
    	    var median int
    	    fmt.Printf("lowerLen: %d \n", lowerLen)
    	    fmt.Printf("upperLen: %d \n", upperLen)
    	    fmt.Println((lowerLen+upperLen) % 2)
    	    if (lowerLen + upperLen) % 2 == 0 {
    	    	if (lowerLen + upperLen) / 2 == lowerLen {
    	    		fmt.Println("a")
    	    		median = lowerHalf.Max()
    	    	} else {
    	    		fmt.Println("b")
    	    		median = upperHalf.Min()
    	    	}
    	    } else {
    	    	if (lowerLen + upperLen + 1) / 2 == lowerLen {
    	    		fmt.Println("c")
    	    		median = lowerHalf.Max()
    	    	} else {
    	    		fmt.Println("d")
    	    		median = upperHalf.Min()
    	    	}
    	    }
    	    sumOfMedians += median
    	    medians = append(medians, median)
	    	fmt.Printf("Median %d\n", median)
    	}
    	fmt.Println("Lower Half")
	    fmt.Println(lowerHalf)
	    fmt.Println("Upper Half")
	    fmt.Println(upperHalf)
	    fmt.Println()
    }
    fmt.Printf("Sum of medians: %d \n", sumOfMedians)
    fmt.Println(medians)
	// fmt.Printf("minimum of upper half: %d\n", higherHalf.Min())
	// for higherHalf.Len() > 0 {
	// 	fmt.Printf("%d ", heap.Pop(higherHalf))
	// }
	// fmt.Println()

	
	// heap.Push(lowerHalf, 2)
	// heap.Push(lowerHalf, 1)
	// fmt.Printf("maximum of lower half: %d\n", lowerHalf.Max())
	// for lowerHalf.Len() > 0 {
	// 	fmt.Printf("%d ", heap.Pop(lowerHalf))
	// }
	// fmt.Println()
}