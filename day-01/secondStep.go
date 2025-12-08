
package main

import (
	"bufio"
	"fmt"
	//"log"
	"os"
	"strings"
	"strconv"
)

func secondStep(){
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
		var goRight bool = false;
		if(strings.Contains(line, "R")){
			goRight = true;
		}


	}
	stringPassword := strconv.Itoa(password)
	fmt.Println(("   the password is: " + stringPassword))
}

/*
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
*/
