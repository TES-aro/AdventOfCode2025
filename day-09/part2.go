package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func part2(scanner *bufio.Scanner){
	xCords := []int{}
	yCords := []int{}
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
		xCords = append(xCords, x)
		yCords = append(yCords, y)
	}
	comX, mx := badCompacting(xCords)
	comY , my:= badCompacting(yCords)

	matrix := drawMatrix(comX, comY, mx, my)

	largest := 0
	// now calculating sizes with the OG
	for i:= 0; i < len(comX) - 1; i++{
		for j := i + 1; j < len(comX); j++{
			x1, x2 := minMax(xCords[i], xCords[j])
			y1, y2 := minMax(yCords[i], yCords[j])
			dx := x2-x1+1
			dy := y2-y1+1
			size := dy*dx
			if size > largest && shapeWithin(matrix,comX[i],comY[i], comX[j],comY[j] ){
				largest = size
			}
		}
	}
	fmt.Printf("the largest shape is: %d",largest)
}



// fuck it, I am brute forcing.
func minMax(a, b int)(i, j int){
	if a <= b{
		return a, b
	}
	return b, a
}

func drawMatrix(xCords, yCords []int, sizeX, sizeY int) [][]bool{
	//using bool as it's the smallest datatype and I just need it to be fiiine
	matrix := [][]bool{}
	for i := 0; i < sizeY; i++{
		line := []bool{}
		for j:= 0; j < sizeX; j++{
			line = append(line, false)
		}
		matrix = append(matrix, line)
	}
	ox := xCords[len(xCords) - 1]
	oy := yCords[len(yCords) - 1]

	for c := 0; c < len(xCords); c++{
		x := xCords[c]
		y := yCords[c]

		if oy == y{
			minX, maxX := minMax(x, ox)
			fmt.Print(minX)
			fmt.Println(maxX)
			for i := minX; i <= maxX; i++{
				matrix[y][i] = true
			}
		}

		if ox == x{
			minY, maxY := minMax(y, oy)
			for j := minY; j <= maxY; j++{
				matrix[j][x] = true
			}
		}
		oy = y
		ox = x
	}
	return matrix
}

func badCompacting(cords []int) ([]int, int){
	sortedCords := []int{}
	for _, v := range cords {
		sortedCords = append(sortedCords, v)
	}
	slices.Sort(sortedCords)
	compressedCords := []int{}
	compressionMap := map[int]int{}
	//fmt.Println(sortedCords) 
	for i, v := range sortedCords {
		_, b := compressionMap[v]
		if b {
			continue
		}
		compressionMap[v] = i
	}
	for _, i := range cords{
		compressedCords = append(compressedCords, compressionMap[i])
	}
	max := compressionMap[sortedCords[len(sortedCords)-1]] + 1
	fmt.Println(max)
	return compressedCords, max
}

func shapeWithin(matrix [][]bool, x1, y1, x2, y2 int) bool{
	x, X := minMax(x1, x2)
	y, Y := minMax(y1, y2)
	// first inside of the shape
	for a := x+1; a < X; a++{
		for b := y+1; b <Y; b++{
			if matrix[b][a] == true{
				return false
			}
		}
	}
	return true
}


