package main

import (
	"bufio"
	"fmt"
	"strings"
)

var nodeMap map[string]*sNode = map[string]*sNode{}

func part2(scanner *bufio.Scanner){
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
	if  dacFirst() {
		countD2F()
		return
	}
	countF2D()
}

func dacFirst() bool {
	dac , b:= nodeMap["dac"]
	if !b {
		fmt.Println("no dac!?")
	}
	fft , b:= nodeMap["fft"]
	if !b {
		fmt.Println("no fft?!")
	}
	d2f := recurseCount(dac, fft)
	Clear()
	f2d := recurseCount(fft, dac)
	Clear()
	fmt.Printf("d2f: %d, f2d: %d\n", d2f, f2d)
	if f2d == 0 {
		return true
	}
	return false
}

func countD2F() int {
	Clear()
	dac := nodeMap["dac"]
	fft := nodeMap["fft"]
	svr , b:= nodeMap["svr"]
	if !b {
	fmt.Println("no svr node!")
	}
	out, b := nodeMap["out"]
	if !b {
		fmt.Println("no out node!")
	}
	first := recurseCount(svr, dac)
	Clear()
	second := recurseCount( dac, fft)
	Clear()
	third := recurseCount(fft, out)
	Clear()
	fmt.Printf("third: %d\n", third)
	fmt.Printf("second: %d\n", second)
	fmt.Printf("first:  %d\n", first)
	fmt.Printf("1*2*3: %d\n",first*second*third)
	return first
}

func countF2D() int {
	dac := nodeMap["dac"]
	fft := nodeMap["fft"]
	svr , b:= nodeMap["svr"]
	if !b {
	fmt.Println("no svr node!")
	}
	out, b := nodeMap["out"]
	if !b {
		fmt.Println("no out node!")
	}
	third := recurseCount(dac, out)
	Clear()
	fmt.Printf("third: %d\n", third)
	second := recurseCount( fft, dac)
	Clear()
	fmt.Printf("second: %d\n", second)
	first := recurseCount(svr, fft)
	fmt.Printf("first: %d\n", first)
	fmt.Printf("1*2*3: %d\n", third*second*first)
	return first
}

func recurseCount(node, goal *sNode)int{
	if node == goal {
		return 1
	}
	if node.visited {
		return node.counter
	}
	node.visited = true
	for _, connectedNode := range node.connectedTo {
		node.counter += recurseCount(connectedNode, goal)
	}
	return node.counter
}

func Clear(){
	for _, node := range nodeMap{
		node.visited = false
		node.counter = 0
	}
}
