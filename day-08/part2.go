package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// eh, it's O(n^2) at least ignoring the sort so let's do it the easy way.


func part2(scanner *bufio.Scanner){
	boxes := []*juncBox{}
	connectedArray := []*juncBox{}
	for scanner.Scan(){
		cordStr := strings.Split(scanner.Text() ,",")
		x , err := strconv.Atoi(cordStr[0])
		if err != nil{
			fmt.Println(cordStr)
			break
		}
		y, err := strconv.Atoi(cordStr[1])
		if err != nil{
			fmt.Println(cordStr)
			break
		}
		z, err := strconv.Atoi(cordStr[2])
		if err != nil{
			fmt.Println(cordStr)
			break
		}
		boxes = append(boxes, &juncBox{x, y,z, map[*juncBox]bool{}, 0})
	}
	// create all possible connections
	i := 0
	connections := []connection{}
	for i < len(boxes) - 1 {
		j := i+1
		for j < len(boxes){
			connections = append(connections, boxes[i].getConnection(boxes[j]))
			j++
		}
		i++
	}
	fmt.Println("sorting connections")
	var lastPair connection
	sort.Slice(connections,
		func(i int, j int) bool { return connections[i].distance < connections[j].distance})
	for i := 0; i < len(connections); i++ {
		pair := connections[i]
		if pair.a.color != 0 && pair.b.color != 0{
			continue
		}
		pair.a.connect(pair.b)
		pair.a.color = 1
		pair.b.color = 1
		connectedArray = append(connectedArray, pair.a)
		lastPair = pair
	}
	fmt.Printf("last pair: %d\n", lastPair.a.x * lastPair.b.x) 

}
