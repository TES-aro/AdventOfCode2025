package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func part1(scanner *bufio.Scanner){
	text := strBuild(scanner)
	operators := text[len(text)-1]
	numberStr := text[0:len(text)-1]
	fmt.Println(operators)
	numbers := toNumbers(numberStr)
	reformat(numbers, operators)
}

func toNumbers(numStrs [][]string) [][]int{
	var numbers [][]int
	for i := range numStrs{
		var innerNum []int
		for j := range numStrs[i]{
			number, err := strconv.Atoi(numStrs[i][j ])
			if err != nil{
				fmt.Println(err)
				continue
			}
			innerNum = append(innerNum, number)
		}
		numbers = append(numbers, innerNum)
	}
	return numbers
}
func reformat(numbers [][]int, operators []string){
	sum := 0
	fmt.Println(operators)
	fmt.Println(len(operators))
	fmt.Print("len of numbers: ")
	fmt.Println(len(numbers))
	fmt.Println(numbers) 
	for o := 0; o < len(operators); o++{
		fmt.Println(o)
		operator := operators[o]
		var newNumbers []int
		for i := range numbers{
			newNumbers = append(newNumbers, numbers[i][o])
		}
		fmt.Println(newNumbers)
		if strings.Contains(operator, "*"){
			sum += multiply(newNumbers)
			continue
		}
		if strings.Contains(operator, "+"){
			sum += summing(newNumbers)
		}
	}
	fmt.Println("Answer:")
	fmt.Println(sum)
}

func summing (numbers []int) int{
	sum := 0
	for i := range numbers{
		sum += numbers[i]
	}
	return sum
}

func multiply (numbers []int) int{
	sum := 1
	for i := range numbers{
		sum = sum * numbers[i]
	}
	return sum
}
