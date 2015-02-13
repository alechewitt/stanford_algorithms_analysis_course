package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
	"math/rand"
    "sort"
)

type Node struct {
	nodeName int
	edgesEndpoints []int
}

func readTextFile(fileName string) []*Node{
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
	return graph
}

func randomContractraction(graph []*Node, channel chan int) int{
	if len(graph) == 2 {
		// Remove self loops
		numberSelfLoops := 0
		for _, endpoint := range graph[0].edgesEndpoints{
            if graph[0].nodeName == endpoint{
                numberSelfLoops += 1
            }
        }
        channel <- (len(graph[0].edgesEndpoints) - numberSelfLoops)
        return (len(graph[0].edgesEndpoints) - numberSelfLoops)
	} else {
		// Select random node index
        nodeIndex := rand.Intn(len(graph)) 

        // Select random node to contract with
        secondNodeValue := graph[nodeIndex].edgesEndpoints[rand.Intn(len(graph[nodeIndex].edgesEndpoints))]

        var newGraph []*Node
        // Rename node to first node name
        mergedNode := new(Node)
        mergedNode.nodeName = graph[nodeIndex].nodeName
        // All nodes that were previously pointing to second node now to point to first node
        for i, node := range graph {
            if i == nodeIndex || node.nodeName == secondNodeValue {
                // Remove second value from our first node
                for _, edgePoint := range graph[i].edgesEndpoints {
                    if edgePoint != secondNodeValue && edgePoint != graph[nodeIndex].nodeName {
                        mergedNode.edgesEndpoints = append(mergedNode.edgesEndpoints, edgePoint)
                    }
                }
            } else {
                // This is not either of the two merged nodes, however we need to check 
                // None of them have an edge pointing to the second node, as will disapear
                newNode := new(Node)
                newNode.nodeName = graph[i].nodeName
                for _, edgePoint := range graph[i].edgesEndpoints {
                    if edgePoint == secondNodeValue{
                        // We need to subsitiute it for the new merged node
                        newNode.edgesEndpoints = append(newNode.edgesEndpoints, mergedNode.nodeName)
                    } else {
                        newNode.edgesEndpoints = append(newNode.edgesEndpoints, edgePoint)
                    }
                }
                newGraph = append(newGraph, newNode)
            }
        }

        // Append merged node to new graph
        newGraph = append(newGraph, mergedNode)
        return randomContractraction(newGraph, channel)
	}

}

func main() {
	graph := readTextFile("testFiles/homeWorkList.txt")

    cutChannel := make(chan int)
    numberAttempts := 10
    var cuts []int
    for i:=0; i < numberAttempts; i++ {
        // Interestingly this method is not quicker!
        // Possibly due to to many task, all being put onto the swap
        // and then a lot of io.
        go randomContractraction(graph, cutChannel)
    }
    for i :=0; i<numberAttempts; i++ {
        numCuts := <- cutChannel
        cuts = append(cuts, numCuts)
    }

    sort.Ints(cuts)
    fmt.Println(cuts)
}

