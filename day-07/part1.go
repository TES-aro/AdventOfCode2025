package main

import (
	"bufio"
	"fmt"
	"utilsAOC"
)

// I was going to do this by recursing but pretty sure line by line is more
// efficient as recursing would require you to go trough the input at least
// twice



func part1(scanner *bufio.Scanner){
	scanner.Scan()
	splitCounter := 0
	beams := utilsAOC.NewSet[int]()
	line := []rune(scanner.Text())
	for i := 0; i < len(line); i++{
		if line[i] == 'S'{
			beams.Add(i)
			break
		}
	}
	for scanner.Scan() {
		line := []rune(scanner.Text())
		if len(line) == 0{
			fmt.Println("empty line – closing scanner\n")
			break
		}
		splitBeam(beams,line,&splitCounter)
	}
	fmt.Printf("beams were split %d times!", splitCounter)
}

func splitBeam(beams *utilsAOC.Set[int], line []rune, counter *int){
	max := len(line) - 1
	values := beams.Values()
 	for v := range values {
		i := values[v]
		if line[i] == '^'{
			*counter++
			beams.Remove(i)
			if i > 1 {
				beams.Add((i - 1))
			}
			if i < (max - 1){
				beams.Add((i + 1))
			}
		}
	}
}

