package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var gifts map[int]*gift = map[int]*gift{}
var areas []area = []area{}

func part1(scanner *bufio.Scanner){
	giftHelper := [][]rune{}
	giftNum := 0
	lineCounter := 0
	for scanner.Scan(){
		line := scanner.Text()
		if lineCounter > 0{
			if line == ""{
				gifts[giftNum] = newGift( giftHelper)
				giftNum++
				giftHelper = [][]rune{}
				lineCounter = 0
				continue
			}
			giftHelper = append(giftHelper, []rune(line))
			continue
		}
		if len(line) == 2 {
			lineCounter++
			continue
		}
		if line == "" {
			continue
		}
		splitLine := strings.Split(line,":")
		dimensions := strings.Split( splitLine[0], "x")
		x , err:= strconv.Atoi(dimensions[0])
		if err != nil {
			fmt.Println("error with cords: ", dimensions[0])
		}
		y, err := strconv.Atoi(dimensions[1])
		if err != nil {
			fmt.Println("error with cords: ", dimensions[1])
		}
		areas = append(areas, newArea(x,y, splitLine[1]))
	}
	counter := 0
	for _, area := range areas {
		if area.greedyCheck() && area.hardCheck() {
			counter++
		}
	}
	fmt.Printf("\n greedy fits: %d\n", counter)

}


// let's start with "graphic" solving and just take the gifts as matrixes
type gift struct {
	shape shape
	size int
}

type shape struct{
	matrix [][]rune
}



type area struct {
	x int
	y int
	giftArray []int
}

func newGift(runes [][]rune) *gift {
	counter := 0
	for _, line := range runes{
		for _, r := range line {
			if r == '#' {
				counter++
			}
		}
	}
	return &gift{shape{runes}, counter}
}

func newArea(x, y int, giftArray string) area{
	area := area{x, y ,[]int{}}
	giftArray = strings.Trim(giftArray, " ")
	for _, num := range strings.Split(giftArray, " "){
		integer , err:= strconv.Atoi(num)
		if err != nil {
			fmt.Println("error with gift array: ", giftArray)
			return area
		}
		area.giftArray = append(area.giftArray, integer)
	}
	return area
}

func (area *area) greedyCheck() bool{
	minGiftSize := 0
	for i, count := range area.giftArray{
		minGiftSize += gifts[i].size*count
	}
	if minGiftSize >= area.x * area.y {
		return false
	}
	return true
}

// I am stypid and *will* brute force this.
func (area *area) hardCheck() bool{
	neededGifts := []*gift{}
	for i, count := range area.giftArray{
		for j := 0; j < count; j++ {
			neededGifts = append(neededGifts, gifts[i])
		}
	}
	for _, gift := range neededGifts{
		fmt.Print(gift.size, " ")
	}
	areaShape := shape{matrix: [][]rune{}}
	for i := 0; i < area.x; i++{
		line := []rune{}
		for j := 0; j < area.y; j++{
			line = append(line, '0')
		}
		areaShape.matrix = append(areaShape.matrix, line)
	}
	return true
}

func (matrix shape) canPlace(shape [][]rune)bool {
 return true
}
