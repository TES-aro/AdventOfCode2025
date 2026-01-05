package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func part2(scanner *bufio.Scanner){
	ranges := []pair{}
	for scanner.Scan(){
		splitLine := strings.Split(scanner.Text(),"-")
		if len(splitLine) < 2{
			continue
		}
		ranges = append(ranges, newRange(splitLine))
	}

	sort.Slice(ranges,func(i, j int) bool { return ranges[j].min > ranges[i].min})
	for _, pair := range ranges{
		fmt.Println(pair) 
	}
	size := len(ranges)
	fmt.Println(size) 
	for true {
		ranges = fuseRanges(ranges)
		newSize := len(ranges)
		fmt.Println(newSize)
		if size == len(ranges) {
			break
		}
		size = len(ranges)
	}
	
	for _, pair := range ranges{
		fmt.Println(pair) 
	}
	count := countAll(ranges)
	fmt.Printf("\nfull count: %d",count)
}

// and now just look at the overlaps
type pair struct{
	min int
	max int
}

func newRange(str []string)pair{
	pair := pair{}
	a, err := strconv.Atoi(str[0])
	if err != nil{
		fmt.Println("something went wrong!")
		return pair
	}
	b, err := strconv.Atoi(str[1])
	if err != nil {
		fmt.Println("fuck")
	}
	x, y := minMax(a,b)
	pair.min = x
	pair.max = y
	return pair
}

func minMax(a, b int)(int,int){
	if a < b {
		return a, b
	}
	return b, a
}

// let's order it to save me some headache. prolly not the most efficient.
func fuseRanges(ranges []pair)[]pair{
	newRanges := []pair{}
	i := 0
	for i < len(ranges) - 1{
		fmt.Printf("comparing %d and %d\n", ranges[i].max, ranges[i+1].min) 
		if ranges[i].max >= ranges[i+1].min{
			fmt.Println("found one") 
			if ranges[i].max >= ranges[i+1].max{
				newRanges = append(newRanges, ranges[i])
				i += 2
				continue
			}
			newPair := pair{ranges[i].min, ranges[i+1].max}
			newRanges = append(newRanges, newPair)
			i += 2
			continue
		}
		newRanges = append(newRanges, ranges[i])
		i++
	}
	if i < len(ranges){
		newRanges = append(newRanges, ranges[i])
	}
	for _, a := range newRanges {
		fmt.Println(a) 
	}

	return newRanges
}

func countAll(ranges []pair)int{
	count := 0
	for _, pair := range ranges{
		count += 1+(pair.max-pair.min)
	}
	return count
}
