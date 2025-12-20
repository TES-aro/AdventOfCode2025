package main

import (
	"fmt"
	"strconv"
)
// quite uggly, but the build in tools aren't quite intuitive.
// could've done better job with sets, I guess
func parsLine(str string) diagram{
	runes := []rune(str)
	lights :=[]bool{}
	buttonsArray := [][]int{}
	buttons := []int{}
	joltage := []int{}
	num := []rune{}
	workingOn := '['
	for _, r := range runes{
		if r == ' '{
			continue
		}
		if r == '.'{
			lights = append(lights, false)
			continue
		}
		if r == '#'{
			lights = append(lights, true)
			continue
		}
		if r == ','{
			x, err := strconv.Atoi(string(num))
			num = []rune{}
			if err != nil{
				fmt.Println( "error parsing in here!")
				fmt.Println(err) 
			}
			if workingOn == '('{
				buttons = append(buttons, x)
				continue
			}
			joltage = append(joltage, x)
			continue
		}
		if r == '(' || r == '{'{
			workingOn = r
			continue
		}
		if r == ')' || r == '}'{
			x, err := strconv.Atoi(string(num))
			num = []rune{}
			if err != nil {
				fmt.Println("error parsing in there!")
				fmt.Println(err)
			}
			if r == ')'{
				buttons = append(buttons, x)
				buttonsArray = append(buttonsArray, buttons)
				buttons = []int{}
				continue
			}
			joltage = append(joltage, x)
			continue
		}
		if r == '[' ||  r == ']'{
			continue
		}
		num = append(num, r)
	}
	/*
	fmt.Println("")
	fmt.Println(lights)
	fmt.Println(buttonsArray)
	fmt.Println(joltage)
	*/
	return diagram{lights, buttonsArray, joltage }

}
