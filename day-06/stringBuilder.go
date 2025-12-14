package main

import (
	"bufio"
	"fmt"
	"strings"
)

func strBuild(scanner *bufio.Scanner) [][]string{
	var text [][]string
	for scanner.Scan(){
		var lineArray []string
		line := scanner.Text()
		// this is an ugly fix. I am  not quite sure why
		// the builder misses the last symbols if line doesn't
		// end in a space.
		lineBuilder := strings.Builder{}
		lineBuilder.WriteString(line)
		lineBuilder.WriteString(" ")
		line = lineBuilder.String()
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


