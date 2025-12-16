package utilsAOC

import (
)

// guess you gotta implement linked list yourself

type iterator[T any] struct{
	currentNode *queueNode[T]
}

type Queue[T any] struct{
	firstNode *queueNode[T]
	lastNode *queueNode[T]
}

type queueNode[T any] struct {
	value T
	previous *queueNode[T]
	next *queueNode[T]
}

func initQueueNode[T any](value T) *queueNode[T]{
	return &queueNode[T]{value, nil,  nil}
}



func (q Queue[T]) Add(value T){
	newNode := initQueueNode[T](value)
	if q.firstNode == nil {
		q.firstNode = newNode
		q.lastNode = newNode
		return
	}
	oldLast := q.lastNode
	newNode.previous = oldLast
	oldLast.next = newNode
	q.lastNode = newNode
}

func (q Queue[T]) NewIterator() iterator[T]{
	return iterator[T]{q.firstNode}
}

func (iter iterator[T]) Remove() {
	prev := iter.currentNode.previous
	next := iter.currentNode.next

	prev.next = next
	next.previous = prev
	iter.currentNode = next
}

func (iter iterator[T]) HasNext() bool{
	if iter.currentNode.next == nil{
		return false
	}
	return true
}

func (iter iterator[T]) Next() T{
	iter.currentNode = iter.currentNode.next
	return iter.currentNode.value
}
