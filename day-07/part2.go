package main

import (
	"bufio"
	"fmt"
)

// I regreted not doing this recursively, but then again implementing
// cache of sort of that would've been more of a hassle


func part2(scanner *bufio.Scanner){
	scanner.Scan()
	splitCounter := 1
	beams := map[int]int{}
	line := []rune(scanner.Text())
	for i := 0; i < len(line); i++{
		if line[i] == 'S'{
			beams[i] = 1
			break
		}
	}
	for scanner.Scan() {
		line := []rune(scanner.Text())
		if len(line) == 0{
			fmt.Println("empty line – closing scanner\n")
			break
		}
		splitBeamExp(&beams,line,&splitCounter)
	}
	fmt.Printf("beams were split %d times!", splitCounter)
}

func splitBeamExp(beams *map[int]int, line []rune, counter *int){
 	for i := range *beams {
		if line[i] == '^'{
			beamCount := (*beams)[i]
			*counter += beamCount
			delete(*beams, i)
			_ , existL:= (*beams)[i-1]
			_, existsR := (*beams)[i+1]
			if existL{
				(*beams)[i-1] += beamCount
			} else {
				(*beams)[i-1] = beamCount
			}
			if existsR{
				(*beams)[i+1] += beamCount
			} else {
				(*beams)[i+1] = beamCount
			}
		}
	}
}

