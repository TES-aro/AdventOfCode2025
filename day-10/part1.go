package main

import (
	"bufio"
	"fmt"
)

func part1(scanner *bufio.Scanner){
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
		answers = append(answers, l.solve())
	}
	fmt.Println(answers)
	sum := 0
	for _, v := range answers{
		sum += v
	}
	fmt.Printf("\n\n  asnwer is: %d",sum)
}

type diagram struct{
	lights []bool
	buttons [][]int
	joltage []int
}

// all start as off, hence we need uneven for those who need to be on and
// and even for off.

func (d *diagram) check(lights []bool) bool{
	if len(d.lights) != len(lights){
		return false
	}
	for i, v := range d.lights{
		if lights[i] != v{
			return false
		}
	}
	return true
}

func pressButton(lights []bool, button []int) []bool{
	newLights := []bool{}
	for _, v := range lights{
		newLights = append(newLights, v)
	}
	for _, i := range button{
		newLights[i] = !newLights[i]
	}
	return newLights
}
// this is kinda funky! let's go level by level this time :v

func (d *diagram) solve()int{
	fmt.Println("the goal is:")
	fmt.Println(d.lights)
	fmt.Println("")
	buttonMap := map[*[]int]int{}
	for i := range d.buttons{
		buttonMap[&d.buttons[i]] = 3
	}
	lightsOff := []bool{}
	for i := 0; i < len(d.lights);i++{
		lightsOff = append(lightsOff, false)
	}
	queue := [][][]bool{}
	layer := [][]bool{}
	layer = append(layer, lightsOff)
	queue = append(queue, layer)
	return goTrough(d)

}

func goTrough(d *diagram) int{
	qu := [][][]bool{}

	lightsOff := []bool{}
	for i := 0; i < len(d.lights);i++{
		lightsOff = append(lightsOff, false)
	}
	layer := [][]bool{}
	layer = append(layer, lightsOff)
	qu = append(qu, layer)
	counter := 0
	nextLayer := [][]bool{}
	currentLayer := layer
	for true{
		counter++
		fmt.Printf("working on layer: %d, with length: %d\n", counter, len(currentLayer))
		for _ , l:= range currentLayer{
			for _, b := range d.buttons{
				newLights := pressButton(l,b)
				if d.check(newLights) {
					fmt.Printf("found the answer at depth: %d\n", counter)
					return counter
				}
				nextLayer = append(nextLayer, newLights)
			}
		}
		if counter+1%10 == 0{
			fmt.Println("something went very wrong with this one!")
			fmt.Println(d.lights)
			return 100
		}
		currentLayer = nextLayer
		nextLayer = [][]bool{}
	}
	return 420
}
