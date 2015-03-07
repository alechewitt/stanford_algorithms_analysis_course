package main

import (
	"fmt")

type IntHeap []int

func (h IntHeap) Len() int { 
	return len(h) 
}

func (h IntHeap) Hello() string {
	return "hello"
}

func main() {
	h := &IntHeap{2, 1, 5}
	fmt.Println("main function")
	fmt.Println(h.Len())
	fmt.Println(h.Hello())
}