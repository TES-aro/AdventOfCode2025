package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1_main(){
	scanner := readFile();
	var strArray [][]string;
	for scanner.Scan(){
		// this time let's read the whole mess into memory
		line := scanner.Text();
		splitLine := strings.Split(line, "")
		strArray = append(strArray, splitLine)
	}
	testArray := strArray[0][0:]
	fmt.Println(len(testArray));
	fmt.Println(strArray[0][4]);
	// center of the 3x3 grid can exist within:
	//xMin := 1;
	xMax := len(strArray[0][0:]) - 1;
	//yMin := 1;
	yMax := len(strArray[0:]) - 1;

	counter := 0;
	for x := 0; x<=xMax; x++{
		fmt.Println(""); 
		for y := 0; y<=yMax; y++{
			if(!isPaper(strArray[x][y])){
				fmt.Print(".") 
				continue;
			}
			if(canMove(x, y, xMax, yMax, strArray)){
				fmt.Print("X")
				counter++
				continue
			}
			fmt.Print("@")
		}
	}

	fmt.Println("\n\n   --- solution ---");
	fmt.Println(counter)
}

func isPaper(str string) bool {
	if(str == "@"){
		return true;
	}
	return false;
}

func canMove(x int, y int, xMax int, yMax int, array [][]string ) bool{
	counter := 0
	for x2 := x -1; x2 <= x+1; x2++{
		if x2 < 0 || x2 > xMax{
			continue;
		}
		for y2 := y -1; y2 <= y+1; y2++{
			if (y2 < 0 || y2 > yMax){
				continue;
			}
			if (y2 == y && x2 == x){
				continue;
			}
			char := array[x2][y2]
			if (isPaper(char)){
				counter++
			}
		}
	}
	if counter < 4{
		//fmt.Print("on line: ")
		//fmt.Print(x)
		//fmt.Print(  "  ")
		//fmt.Println(y)
		return true
	}
	return false

}

func readFile() *bufio.Scanner{
	location := os.Args[1];
	file, err := os.Open(location);
	if( err != nil){
		fmt.Print(err)
	}
	return bufio.NewScanner(file);
}
