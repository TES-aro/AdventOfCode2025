package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	//"strings"
	//"strconv"
)

func part2(scanner *bufio.Scanner){
	var text [][]rune
	for scanner.Scan(){
		substr := scanner.Text()
		text = append(text, []rune(substr)) 
	}
	//operators := text[len(text)-1]
	//numberStr := text[0:len(text)-1]
	formating(text)
}

func formating (text [][]rune){
	fmt.Println(len(text[:]))
	fmt.Println(len(text[1]))
	rotatedText := [][]rune{}
	section := []rune{}
	for y := 0; y < len(text[0]); y++{
		allEmpty := true
		for x := 0; x < len(text); x++{
			selectedRune := text[x][y]
			if string(selectedRune) != " "{
				allEmpty = false
			}
			section = append(section, selectedRune)
		}
		if allEmpty == true{
			rotatedText = append(rotatedText, section)
			section = []rune{}
		}
	}
	rotatedText = append(rotatedText, section) 
	fmt.Println("----roteted text:") 
	for r := range rotatedText{
		fmt.Println(string(rotatedText[r]))
	}
	parseRunes(rotatedText) 
}

func parseRunes(text [][]rune){
	fmt.Printf("--- size of parseRune input: %d ---\n", len(text)) 
	snailArray := []snailMath{}
	for i := range text{
		snailObject := snailMath{}
		group := text[i]
		builder := strings.Builder{}
		for j := range group{
			symbol := string(group[j])
			fmt.Print(symbol + " ") 
			if symbol == "+" || symbol == "*"{
				snailObject.operator = symbol
				str := builder.String()
				builder = strings.Builder{}
				if str != ""{
					num , err := strconv.Atoi(str)
					if err != nil {
						fmt.Println(err) 
					}
					snailObject.numbers = append(snailObject.numbers, num)
				}
				continue
			}
			if isNum(symbol) == true{
				builder.WriteString(symbol)
				continue
			}
			str := builder.String()
			builder = strings.Builder{}
			if str != ""{
				num , err := strconv.Atoi(str)
				if err != nil {
					fmt.Println(err) 
				}
				snailObject.numbers = append(snailObject.numbers, num)
			}
		}
		snailArray = append(snailArray, snailObject)
	}
	for s := range snailArray{
		fmt.Printf("object num:%d\nopreator: %s\n", s, snailArray[s].operator)
		fmt.Println(snailArray[s].numbers)
	}
	doMath(snailArray) 
}

func doMath(snailArray []snailMath){
	sum := 0
	fmt.Print("\n\n\ndoMath:\n")
	fmt.Println(len(snailArray)) 
	for i := range snailArray{
		fmt.Print("snail:")
		snail := snailArray[i]
		fmt.Println(snail.numbers)
		fmt.Println(snail.operator) 
		if snail.operator == "*"{
			subsum := 1
			for s := range snail.numbers{
				snailNum := snail.numbers[s]
				fmt.Println(snailNum)
				subsum = subsum * snailNum 
			}
			fmt.Println(subsum) 
			sum += subsum
			continue
		}
		subsum := 0
		for s := range snail.numbers{
			snailNum := snail.numbers[s]
			subsum += snailNum
		}
		fmt.Println("subsum:") 
		fmt.Println(subsum)
		sum += subsum
	}
	fmt.Printf("\n\n  answer is:\n    %d \n",sum)
}
func isNum(symbol string) bool{
	_, err := strconv.Atoi(symbol)
	if err != nil{
		return false
	}
	return true
}
type snailMath struct{
	operator string
	numbers []int
}



func numBuild(numstr []string) [][]string{
	var text [][]string
	for i := range numstr{
		var lineArray []string
		line := numstr[i]
		fmt.Println("line:")
		fmt.Println(line) 
		// this is an ugly fix. I am  not quite sure why
		// the builder misses the last symbols if line doesn't
		// end in a space.
		//lineBuilder := strings.Builder{}
		//lineBuilder.WriteString(line)
		//lineBuilder.WriteString(" ")
		//line = lineBuilder.String()
		builder := strings.Builder{}
		for _, c := range []rune(line){
			if string(c) != " "{
				builder.WriteString(string(c))
				continue
			}
			str := builder.String();
			if str != ""{
				lineArray = append(lineArray, str)
			}
			builder = strings.Builder{}
		}
		text = append(text, lineArray)
	}
	fmt.Println("text:") 
	fmt.Println(text)
	return text
}
