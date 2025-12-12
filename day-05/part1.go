package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
// guess I'll implement linked list meself
type LinkedList struct{
	head  *Node
	tail *Node
	current *Node
}

type Node struct{
	data int
	next *Node
	prev *Node
}

// InsertAtEnd inserts a new node at the end of the list
func (ll *LinkedList) Add(data int) {
    newNode := &Node{data: data}
    if ll.head == nil {
	    fmt.Println("Started a new list!")
        // If the list is empty, set both head and tail to the new node
        ll.head = newNode
        ll.tail = newNode
        ll.current = newNode
    } else {
        ll.tail.next = newNode
        newNode.prev = ll.tail
        ll.tail = newNode
    }
}

func (ll *LinkedList) Start() int{
	ll.current = ll.head
	return ll.current.data
}
func (ll *LinkedList) Next() (int, string){
	if ll.current == nil{
		return 0, "no next found"
	}
	val := ll.current.data
	ll.current = ll.current.next
	return val, ""
}
func (ll *LinkedList) Delete(){

}

func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}


/*
func (ll LinkedList) Add(val int){
	if(ll.end == nil){
		newNode := &listNode{value: val}
		ll.startNode = newNode
		ll.endNode = newNode
		return
	}
	newNode := &listNode{val, nil, ll.endNode}
	endNode := ll.endNode
	endNode.nextNode = newNode
	ll.endNode = newNode
}

func (ln listNode) HasNext() bool{
	if ln.nextNode != nil{
		return true
	}
	return false
}

func (ln listNode) Remove(){
	prev := ln.prevNode
	next := ln.nextNode
	prev.nextNode = next
	next.prevNode = prev
}
*/



// back to code proper

func part1(scanner *bufio.Scanner){
	var ranges []IDrange
	items := *new(LinkedList)
	llCounter := 0;
	for scanner.Scan(){
		line := scanner.Text();
		if line == "" {
			fmt.Println("blank line detected")
			for scanner.Scan(){
				line := scanner.Text();
				items.Add( toInt(line))
				llCounter++
			}
			break
		}
		ranges = addIDrange(ranges, line)
	}
	fmt.Println(llCounter)
	compare(ranges, items)
}

func compare(ranges []IDrange, items LinkedList){
	freshCount := 0
	for i:= 0; i < len(ranges); i++{
		Range := ranges[i]
		//fmt.Println("now handling range")
		//fmt.Println(idRange)
		node := items.head

		for node != nil{
			val := node.data
			if val <= Range.end && val >= Range.start{
				freshCount++
				nextNode := node.next
				prevNode := node.prev

				if node == items.head {
					items.head = node.next
					items.head.prev = nil
				}

				if node == items.tail {
					items.tail = node.prev
					node.prev.next = nil
				}
				if(nextNode != nil){
					nextNode.prev = prevNode
				}
				if(prevNode != nil){
					prevNode.next = nextNode
				}
			}
			node = node.next
		}
	}
	fmt.Println("answer is:")
	fmt.Println(freshCount)
}

type IDrange struct {
	start int
	end int
}
func toInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("issue converting '" + str + "' to int")
		os.Exit(42)
	}
	return num
}
func addIDrange(array []IDrange, valuePair string) []IDrange{
	helperPair := strings.Split(valuePair, "-")
	var pair = IDrange{toInt(helperPair[0]), toInt(helperPair[1])}
	array = append(array, pair)
	return array
}


