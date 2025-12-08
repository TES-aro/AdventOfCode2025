package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := readFile();
	sumJoltage := 0;
	for scanner.Scan(){
		line := scanner.Text();
		bestJoltage := toNum( recurse(line, 12));
		fmt.Print(bestJoltage);
		fmt.Print("  ")
		sumJoltage += bestJoltage
	}
	prnt("\n\n   ---sum---")
	fmt.Println(sumJoltage)
}
// let's brute force. O(n^3) here we coooome
// recursion sounds smart for part 2
// very inefficiense as it *will* redo same bit multiple time

func compare (str1 string, str2 string) bool {
	num1 := toNum(str1);
	num2 := toNum(str2);
	if num1 > num2{
		return true
	}
	return false
}
func recurse (str string, levelsLeft int) string {
	levelsLeft--
	strEnd := len(str) - levelsLeft;
	biggestNum :=  str[0:1];
	biggestIndex := 1;
	for i := 1; i <= strEnd; i++{
		num := str[i-1:i];
		if (compare(num, biggestNum)){
			biggestNum = num
			biggestIndex = i;
		}
	}
	if levelsLeft > 0 {
		recurseString := recurse(str[biggestIndex:], levelsLeft)
		return biggestNum+recurseString;
	}

	return biggestNum;
}
func compare_p1(str string) int{
	firstCharPair := str[0:2]
	bestNumPair := toNum(firstCharPair);
	for i := 1; i < len(str); i++{
		firstChar := str[i-1:i]
		for j := i; j <= len(str); j++{
			if(i == j){
				continue;
			}
			compChar := str[j-1:j]
			newCharPair := firstChar + compChar
			newNumPair := toNum(newCharPair)
			if(newNumPair > bestNumPair){
				bestNumPair = newNumPair;
			}
		}
	}
	return bestNumPair
}

func toNum(str string) int{
	newNumt , err:= strconv.Atoi(str);
	if(err != nil){
		prntE(err);
	}
	return newNumt;
}

func prnt(str string){
	fmt.Println(str) 
}

func prntE(err error){
	fmt.Println(err) 
}
func readFile() *bufio.Scanner{
	location := os.Args[1];
	file, err := os.Open(location);
	if( err != nil){
		fmt.Print(err)
	}
	return bufio.NewScanner(file);
}
