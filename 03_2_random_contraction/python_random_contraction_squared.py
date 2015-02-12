"""
Super inefficent algorithm for calculating the minimum cut of
a graph. A Graph is represented in the following form:
    [ [1, [2, 3, 4]], [2, [1, 3, 4] ] ]
    where graph[n][0] == a node
    and graph[n][1] == the nodes its connected to via an edge
"""



import random

def readInInputFile(fileName):
    with open(fileName, "r") as inputFile:
        graph = []
        for line in inputFile:
            tempArray = [int(i) for i in line.split()]
            graph.append([tempArray.pop(0), tempArray])

    return graph

def contractionAlgorithm(aGraph):
    # print "\nGraph to contract"
    # print aGraph
    if len(aGraph) == 2:
        # Remove self loops 
        numberSelfLoops = 0
        for i in aGraph[0][1]:
            if aGraph[0][0] == i:
                numberSelfLoops += 1
        return len(aGraph[0][1]) - numberSelfLoops
    else:
        # Select random node index
        nodeIndex = random.randrange(0, len(aGraph))

        # Select random node to contract with
        secondNodeValue = aGraph[nodeIndex][1][random.randrange(0, len(aGraph[nodeIndex][1]))]

        newGraph = []
        # Rename node to first node name
        mergedNode = [aGraph[nodeIndex][0], []]
        # All nodes that were previously pointing to second node now to point to first node
        for i, value in enumerate(aGraph):
            if i == nodeIndex or value[0] == secondNodeValue:
                # Remove second value from our first node
                for edgePoint in aGraph[i][1]:
                    if edgePoint != secondNodeValue and edgePoint != aGraph[nodeIndex][0]:
                        mergedNode[1].append(edgePoint)

            else:
                # This is not either of the two merged nodes, however we need to check 
                # None of them have an edge pointing to the second node, as will disapear
                newValue = [aGraph[i][0], []];
                for edgePoint in aGraph[i][1]:
                    if edgePoint == secondNodeValue:
                        # We need to subsitiute it for the new merged node
                        newValue[1].append(mergedNode[0])
                    else:
                        newValue[1].append(edgePoint)
                newGraph.append(newValue)

        # Append merged node to new graph
        newGraph.append(mergedNode)
        # print "New Graph = " 
        # print newGraph
        return contractionAlgorithm(newGraph)

testGraph = readInInputFile("testFiles/homeWorkList.txt")

minCuts = []
for i in range(len(testGraph) ** 2):
    minCuts.append(contractionAlgorithm(testGraph))

print "\nMinumum Cut = " + str(min(minCuts))


