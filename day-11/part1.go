package main

import (
	"bufio"
	"fmt"
	"strings"
)

func part1(scanner *bufio.Scanner){
	nodeMap := make(map[string]*sNode)
	for scanner.Scan(){
		line := scanner.Text()
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, ":")
		connectsTo := strings.Split(splitLine[1], " ")
		mainNode, b := nodeMap[splitLine[0]]
		if !b {
			mainNode = CreateNode(splitLine[0])
			nodeMap[splitLine[0]] = mainNode
		}
		for _, nodeName := range connectsTo{
			nodeName = strings.ReplaceAll(nodeName, " ", "")
			if nodeName == ""{
				continue
			}
			node, b := nodeMap[nodeName]
			if !b {
				node = CreateNode(nodeName)
				nodeMap[nodeName] = node
			}
			mainNode.connectedTo = append(mainNode.connectedTo, node)
		}
	}
	start , exists:= nodeMap["you"]
	if !exists {
		fmt.Println("Someting went wrong and start node is not found")
	}
	counter := traverseGraph(start)
	fmt.Printf("final count is %d",counter)
}

func traverseGraph(node *sNode) int{
	if node.name == "out"{
		return 1
	}
	counter := 0
	for _, connectedNode := range node.connectedTo{
		counter += traverseGraph( connectedNode)
	}
	return counter
}

type sNode struct{
	name string
	connectedTo []*sNode
	visited bool
	counter int
}

func CreateNode(name string) *sNode{
	return &sNode{name,[]*sNode{}, false, 0}
}

func (parent *sNode) ConnectTo(node *sNode){
	parent.connectedTo = append(parent.connectedTo, node)
}
