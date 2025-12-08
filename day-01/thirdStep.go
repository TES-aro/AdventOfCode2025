
package main

import (
	"bufio"
	"fmt"
	//"log"
	"os"
	"strings"
	"strconv"
)

func main(){
	var path string = os.Args[1];
	file, err := os.Open(path)
	if(err != nil){
		fmt.Println(("no file at location " + path))
	}
	defer file.Close()
	scanner := bufio.NewScanner((file))

	var zeroCount int = 0;
	var position int = 50;
	for(scanner.Scan()){
		line := scanner.Text()
		var goRight bool = strings.Contains(line, "R")

		steps,err := strconv.Atoi(line[1:])
		if(err != nil){
			fmt.Println("---ERROR!---")
			fmt.Println((err))
		}

		// this is extremely inefficient way of handlin it as it's O(n) where
		// n == sum of each movement
		// at this point I am doing it to sanitycheck.

		i := 0;

		if(goRight){
			for ( i < steps){
				position += 1;
				if(position > 99){
					position = 0
				}
				if(position == 0){
					zeroCount++;
				}
				i++
			}
		} else {
			for ( i < steps){
				position -= 1;
				if(position == 0){
					zeroCount++;
				}
				if(position < 0){
					position = 99;
				}
				i++
			}
		}
	}
	stringPassword := strconv.Itoa(zeroCount)
	fmt.Println(("   the password is: " + stringPassword))
}
