package main

import (
	"Util/queue"
	"fmt"
)

func main(){
	test := queue.Queue[int]{}
	test.Add(1)
	test.Add(2)
	node := test.GetFirst()
	fmt.Print(node.Value)

}
