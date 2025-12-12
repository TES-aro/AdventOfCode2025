package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println("hi")
	part1(readFile())
}

func readFile() *bufio.Scanner{
	if len(os.Args) < 2 {
		fmt.Println("  ERROR:\n  file location required")
		os.Exit(2)
	}
	location := os.Args[1];
	file, err := os.Open(location);
	if( err != nil){
		fmt.Println("  ERROR:\n  no file found at: " + os.Args[1])
		os.Exit(1)
	}
	return bufio.NewScanner(file);
}
