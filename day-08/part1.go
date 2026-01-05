package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// eh, it's O(n^2) at least ignoring the sort so let's do it the easy way.

const connectionCount int = 1000

func part1(scanner *bufio.Scanner){
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
	sort.Slice(connections,
		func(i int, j int) bool { return connections[i].distance < connections[j].distance})
	for i := 0; i < connectionCount; i++ {
		pair := connections[i]
		pair.a.connect(pair.b)
		connectedArray = append(connectedArray, pair.a)
	}
	colorCounter(connectedArray)
}

type juncBox struct {
	x int
	y int
	z int
	connectedTo map[*juncBox]bool
	color int
}

type connection struct {
	distance float64
	a *juncBox
	b *juncBox
}

func  distance(a, b *juncBox) float64{
	distance := math.Sqrt(math.Pow(float64(a.x - b.x), 2) +
		math.Pow(float64(a.y - b.y), 2) + math.Pow(float64(a.z-b.z),2))
	return distance
}

func (a *juncBox) getConnection(b *juncBox) connection{
	return connection{distance(a, b), a, b}
}

func (a *juncBox) connect(b *juncBox){
	a.connectedTo[b] = true
	b.connectedTo[a] = true
}

func recurseColor(box *juncBox, color int)int{
	counter := 0
	if box.color == color{
		return counter
	}
	box.color = color
	for  connectedBox := range box.connectedTo{
		counter += recurseColor(connectedBox, color)
	}
	return counter + 1
}

func colorCounter(boxArray []*juncBox)int{
	color := 1
	counts := []int{}
	fmt.Printf("size of juncBox: %d\n",len(boxArray))
	for _, box := range boxArray{
		if box.color != 0 {
			continue
		}
		newCount := recurseColor(box, color)
		fmt.Println(newCount)
		counts = append(counts, newCount)
	}
	sort.Slice(counts,  func(i int, j int) bool{return counts[i]>counts[j]})
	multiplied := 1
	for i := 0; i < 3; i++ {
		multiplied = multiplied * counts[i]
	}
	fmt.Printf("multiplied size is: %d\n", multiplied)
	return multiplied
}
