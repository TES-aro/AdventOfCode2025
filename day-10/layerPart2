package main

import (
	"bufio"
	"fmt"
)

func part2(scanner *bufio.Scanner){
	instructions := []*diagram{}
	for scanner.Scan(){
		line := scanner.Text()
		if line == ""{
			continue
		}
		diagram := parsLine(line) 
		instructions = append(instructions, &diagram)
	}

	answers := []int{}
	for _, l := range instructions{
		answers = append(answers, l.solveJolt())
	}
	fmt.Println(answers)
	sum := 0
	for _, v := range answers{
		sum += v
	}
	fmt.Printf("\n\n  asnwer is: %d",sum)
}


// all start as off, hence we need uneven for those who need to be on and
// and even for off.


func (d *diagram) checkJolt(jolts []int) int{
	if len(d.joltage) != len(jolts){
		return -999
	}
	for i, v := range d.joltage{
		if jolts[i] > v{
			return -1
		}
		if jolts[i] < v {
			return 0
		}
	}
	return 1
}

func pressJoltButton(joltage []int, button []int) []int{
	newJolt := []int{}
	for _, v := range joltage{
		newJolt = append(newJolt, v)
	}
	for _, i := range button{
		newJolt[i] += 1
	}
	return newJolt
}
// this is kinda funky! let's go level by level this time :v

func (d *diagram) solveJolt()int{
	fmt.Println("the goal is:")
	fmt.Println(d.lights)
	fmt.Println("")
	lightsOff := []bool{}
	for i := 0; i < len(d.lights);i++{
		lightsOff = append(lightsOff, false)
	}
	queue := [][][]bool{}
	layer := [][]bool{}
	layer = append(layer, lightsOff)
	queue = append(queue, layer)
	return goTroughJolt(d)

}

func goTroughJolt(d *diagram) int{
	lightsOff := []int{}
	for i := 0; i < len(d.lights);i++{
		lightsOff = append(lightsOff, 0)
	}
	layer := [][]int{}
	layer = append(layer, lightsOff)
	counter := 0
	nextLayer := [][]int{}
	currentLayer := layer
	for true{
		counter++
		fmt.Printf("working on layer: %d, with length: %d\n", counter, len(currentLayer))
		for _ , l:= range currentLayer{
			for _, b := range d.buttons{
				newLights := pressJoltButton(l,b)
				check := d.checkJolt(newLights)
				if check == 1{
					fmt.Printf("found the answer at depth: %d\n", counter)
					return counter
				}
				if check == 0{
					nextLayer = append(nextLayer, newLights)
				}
			}
		}
		if counter > 500{
			fmt.Println("something went very wrong with this one!")
			fmt.Println(d.lights)
			return 666
		}
		currentLayer = nextLayer
		nextLayer = [][]int{}
	}
	return 420
}
