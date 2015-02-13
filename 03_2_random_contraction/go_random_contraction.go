package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

type Node struct {
	nodeName int
	edgesEndpoints []int
}

func readTextFile(fileName string){
	var graph []*Node
	inFile, _ := os.Open(fileName)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines) 

	for scanner.Scan() {
		node := new(Node)
		stringData := strings.Fields(scanner.Text())
		for i, value := range stringData	{
			intNumber, _ := strconv.Atoi(value)
			if i == 0 {
				node.nodeName = intNumber
			} else {
				node.edgesEndpoints = append(node.edgesEndpoints, intNumber)
			}
		}
		graph = append(graph, node)
	}
	fmt.Println(graph)
}

func main() {
	readTextFile("testFiles/fileRead2.txt")	
}

