package main

import (
	"bufio"
	"fmt"
	"maps"
	"math"
	"sort"
	"strconv"
	"strings"
)

// eh, it's O(n^2) at least ignoring the sort so let's do it the easy way.

const connectionCount int = 20

func part1(scanner *bufio.Scanner){
	boxes := []juncBox{}
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
		boxes = append(boxes, juncBox{x, y,z, map[*juncBox]bool{}, 0})
	}
	// calculate each connection
	connectionLengths := getDistances(boxes)
	sort.Slice(connectionLengths, func(i, j int) bool {
		return connectionLengths[i].distance < connectionLengths[j].distance
	})
	for i := 0; i < connectionCount; i++{
		connection := &connectionLengths[i]
		connection.connect()
	}
	circuits := findconnected(boxes)
	sort.Slice(circuits,func(i, j int) bool {
		return circuits[i].size > circuits[j].size
	})
	// last getting the three largest
	answer := 1
	for i := 0; i < 3; i++{
		fmt.Println(circuits[i])
		answer = answer * circuits[i].size
	}
	fmt.Printf("the answer is: %d", answer) 
}

type juncBox struct{
	x int
	y int
	z int
	connected map[*juncBox]bool
	// group 0 means not connected.
	group int
}

func (jb *juncBox) distance(target *juncBox) float64{
	x := jb.x - target.x
	y := jb.y - target.y
	z := jb.z - target.z
	sum := x*x + y*y + z*z
	return math.Pow(float64(sum), 0.5)
}

type boxConnection struct{
	boxes map[*juncBox]bool
	distance float64
}

func (jb *juncBox) getBoxConnection(target *juncBox) boxConnection{
	distance := jb.distance(target)
	boxes := map[*juncBox]bool{}
	boxes[jb] = true
	boxes[target] = true
	return boxConnection{boxes, distance}
}

func getDistances(boxes []juncBox) []boxConnection{
	connections := []boxConnection{}
	for i := range boxes{
		box := &boxes[i]
		for j := i+1; j < len(boxes); j++{
			targetBox := &boxes[j]
			connections = append(connections, box.getBoxConnection(targetBox))
		}
	}
	return connections
}

func sortConnections(connections []boxConnection) []boxConnection{
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})
	return connections
}

func (con *boxConnection) connect(){
	boxes := []*juncBox{}
	for a:= range con.boxes{
		boxes = append(boxes, a)
	}
	jb := boxes[0]
	target := boxes[1]
	if jb.connected[target] != true{
		jb.connected[target] = true
		target.connected[jb] = true
	}
}

type circuitSize struct {
	number int
	size int
}


func findconnected(boxes []juncBox) []circuitSize{
	groups := []circuitSize{}
	groupNum := 0
	for _ , box:= range boxes{
		if box.group != 0 {
			continue
		}
		groupNum++
		count := box.color(groupNum)
		groups = append(groups, circuitSize{groupNum, count})
	}
	return groups
}

func (jb *juncBox) color(num int) int{
	if jb.group != 0{
		return 0
	}
	count := 1
	jb.group = num
	box := *jb
	for connection := range maps.Keys(box.connected){
		count += connection.color(num)
	}
	return count
}
