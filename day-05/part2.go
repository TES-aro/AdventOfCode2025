
package main

import (
	"bufio"
	"fmt"
	"strings"
)
// guess I'll implement linked list meself
type IDLinkedList struct{
	head  *IDNode
	tail *IDNode
	current *IDNode
}

type IDNode struct{
	data IDrange
	next *IDNode
	prev *IDNode
}

// InsertAtEnd inserts a new node at the end of the list
func (node *IDNode) Delete(list *IDLinkedList){
	prev := node.prev
	next := node.next
	if node == list.head {
		list.head = node.next
		list.head.prev = nil
	}

	if node == list.tail {
		list.tail = node.prev
		node.prev.next = nil
	}
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
}
func (ll *IDLinkedList) Add(data IDrange) {
    newNode := &IDNode{data: data}
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

// back to code proper

func part2(scanner *bufio.Scanner){
	ranges := *new(IDLinkedList)
	llCounter := 0;
	for scanner.Scan(){
		line := scanner.Text();
		if line == "" {
			fmt.Println("blank line detected")
			break
		}
		llCounter++
		hPair := strings.Split(line, "-")
		var pair = IDrange{toInt(hPair[0]), toInt(hPair[1])}
		ranges.Add(pair)
	}
	fmt.Println(llCounter)
	iterate(ranges)
}

func fullOverlap (first IDrange, second IDrange) bool{
	fStart := first.start
	fEnd := first.end
	sStart := second.start
	sEnd := second.end

	if fEnd < sEnd && fStart > sStart{
		return true
	}
	return false
}

func overlapStart(first IDrange, second IDrange) bool{
	fStart := first.start
	sStart := second.start
	sEnd := second.end

	if fStart <= sEnd && fStart >= sStart{
		return true
	}
	return false
}


func overlapEnd(first IDrange, second IDrange) bool{
	fStart := first.start
	fEnd := first.end
	sStart := second.start
	sEnd := second.end

	if fS§
}	
func overlap(first IDrange, second IDrange) bool{
	fStart := first.start
	fEnd := first.end
	sStart := second.start
	sEnd := second.end
	if fEnd < sStart || sEnd < fStart{
		return false
	}
	// delete these.
	if fEnd < sEnd && fStart > sStart{
		return true
	}
	if fEnd >= sStart && fStart < sStart{
		return true
	}
	return false
}

func iterate(rangeList IDLinkedList){
	for true {
		delCounter := 0
		nodePrime := rangeList.head
		for nodePrime != nil {
			pPair := nodePrime.data
			pStart := pPair.start
			pEnd := pPair.end
			nodeSecond := rangeList.head
			for nodeSecond != nil{
				if nodeSecond == nodePrime{
					nodeSecond = nodeSecond.next
					continue
				}
				sPair := nodeSecond.data
				sStart := sPair.start
				sEnd := sPair.end
				if sEnd <= pEnd && sStart >= pStart {
					delCounter++
					nodeSecond.Delete(&rangeList) 
				}
				if (sStart <= pStart){
					if(sEnd >= pStart){
						nodePrime.data.start = sStart
						if(sEnd >= pEnd){
							nodePrime.data.end = sEnd
						}
						//deleting manually yay
						delCounter++
						nodeSecond.Delete(&rangeList)
					}
				}
				if (sEnd >= pEnd && sStart <= pEnd){
					nodePrime.data.end = sEnd;
					delCounter++
					nodeSecond.Delete(&rangeList)
				}
 				nodeSecond = nodeSecond.next
			}
			nodePrime = nodePrime.next
		}
		fmt.Println(delCounter) 
		if delCounter == 0 {
			break
		}
	}
	//once more!
	node := rangeList.head
	counter := 0
	ingredientCounter := 0
	for node != nil {
		counter++
		fmt.Print(node.data.start)
		fmt.Print(" - ")
		fmt.Println(node.data.end) 
		ingredientCounter += (node.data.end - node.data.start + 1)
		node = node.next
	}
	fmt.Println("new list size:")
	fmt.Println(counter)
	fmt.Println("Answer is:")
	fmt.Println(ingredientCounter) 
}
