package main 

import (
	"container/heap"
	"fmt"
)

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
	h := &MinHeap{[]int{2, 1, 5}}
	heap.Init(h)
	heap.Push(h, 3)
	// fmt.Printf("minimum: %d\n", (*h)[0])
	fmt.Printf("minimum: %d\n", h.Min())
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()

	h1 := &MaxHeap{[]int{2, 1, 5}}
	heap.Init(h1)
	heap.Push(h1, 3)
	fmt.Printf("maximum: %d\n", h1.Max())
	for h1.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h1))
	}
	fmt.Println()
}
