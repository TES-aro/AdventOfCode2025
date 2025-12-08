package main

import "fmt"

func print_test_main (){
	testString := "1234567"
	prnt(testString[0:1])
	prnt(testString[1:2])
	prnt(testString[3:5])
	prnt(testString[0:0])
	str := "1"
	fmt.Println((len(str))) 
}

func prnt(str string){
	fmt.Println(str) 
}
