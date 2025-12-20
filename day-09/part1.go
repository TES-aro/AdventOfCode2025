package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner){
	cords := [][]int{}
	for scanner.Scan(){
		line := scanner.Text()
		if line == ""{
			break
		}
		pair := strings.Split(line ,",")
		x , err := strconv.Atoi(pair[0])
		if err != nil {
			break
		}
		y, err := strconv.Atoi(pair[1])
		if err != nil{
			break
		}
		cords = append(cords, []int{x,y})
	}

	largest := 0

	for i := 0; i < len(cords) - 1; i++{
		prime := cords[i]
		for j := i + 1; j < len(cords); j++{
			comp := cords[j]
			x := prime[0] - comp[0]
			if x < 0 {
				x = x * -1
			}
			y := prime[1] - comp[1]
			if y < 0 {
				y = y * -1
			}
			size := (x+1) * (y+1)
			if size > largest{
				largest = size
			}
		}
	}
	fmt.Printf("largest rectangle is: %d", largest)
}
