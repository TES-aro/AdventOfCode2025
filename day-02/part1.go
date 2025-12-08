package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"strings"
)
func part1(){
	if len(os.Args) < 2{
		fmt.Println("no file given.")
		return
	}
	path := os.Args[1]
	file, err := os.Open(path)
	if(err != nil){
		fmt.Println((err));
		return;
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	imput := scanner.Text();
	fmt.Println("the answer is:")
	fmt.Println(breakDown(imput));
	//rangeArray := strings.Split(file.,",")
}

func breakDown(imputs string) int{
	imputArray := strings.Split((imputs), ",")
	returnSum := 0;
	for i := 0; i < len(imputArray); i++ {
		valuePair := strings.Split(imputArray[i], "-");
		start, err := strconv.Atoi(valuePair[0])
		if( err != nil ){
			fmt.Println("\n\nimproper value: " + valuePair[0]);
			fmt.Println((err));
			break;
		}
		end, err := strconv.Atoi(valuePair[1]);
		if( err != nil ){
			fmt.Println("\n\nimproper value: " + valuePair[1]);
			fmt.Println((err));
			break;
		}
		returnSum += solver(start, end);
	}
	return returnSum
}

func solver (start int, end int) int {
	returnSum := 0;
	var repeatingArray []int
	ID := start;
	for (ID <= end){
		if(isRepeating(ID)) {
			repeatingArray = append(repeatingArray, ID);
		}
		ID++;
	}
	fmt.Println(repeatingArray)
	for i := 0; i < len(repeatingArray); i++{
		returnSum += repeatingArray[i]
	}
	return returnSum;
}

func toString(num int) string{
	return strconv.Itoa(num)
}

func isRepeating(ID int) bool {
	str := toString(ID);
	strLength := len(str);
	if(strLength % 2 != 0){
		return false;
	}
	strSize := strLength / 2;
	str1 := str[0:strSize];
	str2 := str[strSize:strLength]
	if (str1 == str2){
		return true;
	}
	return false;
}


func isRepeatingOld(ID int) bool {
	str := toString(ID);
	//fmt.Println("\ntesting now: " + str)
	strLength := len(str);
	upperLimit := strLength/2
	////fmt.Println("upper limit: " + toString(upperLimit) + " strLength: " + toString(strLength) )
	for i := 1; i <= upperLimit; i++ {
		if strLength % i != 0 {
			//fmt.Println("something broke with i: " + toString(i) + " str length: " + toString(strLength) )
			continue;
		}
		j := i
		firstString := str[0:j]
		//fmt.Println("new string: " + firstString)
		for j <= (strLength - i) {
			newString := str[j:(j+i)]
			j += i;

			if newString != firstString {
				//fmt.Println(" _-_-_ does not repeat: " + newString)
				break;
			}
			//fmt.Println("--- repeating: " + newString)
			if(j == strLength){
				//fmt.Println("|||| full repetition with: " + firstString + " and " + str)
				return true;
			}
		}
	}
	return false;
}
