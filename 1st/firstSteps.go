
package main

import (
	"bufio"
	"fmt"
	//"log"
	"os"
	"strings"
	"strconv"
)

func firstStep(){
	fmt.Println(len(os.Args), os.Args)
	var path string = os.Args[1];
	file, err := os.Open(path)
	if(err != nil){
		fmt.Println(("no file at location " + path))
	}
	defer file.Close()
	scanner := bufio.NewScanner((file))

	var password int = 0;
	var position int = 50;
	for(scanner.Scan()){
		line := scanner.Text()
		if(strings.Contains(line, "R")){
			line = strings.Trim(line, "R")
			move, err := strconv.Atoi(line)
			if(err != nil){
				fmt.Println((err))
			}
			position = moveRight(position, move)
		} else {
			line = strings.Trim(line, "L")
			move, err := strconv.Atoi(line)
			if(err != nil){
				fmt.Println((err))
			}
			position = moveLeft(position, move)
		}
		if(position == 0){
			fmt.Println((strconv.Itoa(password) +" with movement:" + scanner.Text())) 
			password++
		}
	}
	stringPassword := strconv.Itoa(password)
	fmt.Println(("   the password is: " + stringPassword))
}

func readLine() {
	
}
func moveRight(start int, steps int) int {
	position:= start + steps
	for position > 99 {
		position -= 100
	}
	return position
}

func moveLeft(start int, step int) int {
	position := start - step
	for position < 0 {
		position += 100
	}
	return position
}
