package linkedlist

import (
	"errors"
)

type Node struct {
	Value          interface{}
	NextNode     *Node
	PreviousNode *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

var ErrEmptyList = errors.New("list has length zero")

func NewList(args ...interface{}) *List {
	newList := &List{}
	for _, value := range args {
		newList.Push(value)
	}
	return newList
}

func (n *Node) Next() *Node {
	return n.NextNode
}

func (n *Node) Prev() *Node {
	return n.PreviousNode
}

func (l *List) Unshift(v interface{}) {
	l.head = &Node{Value: v, NextNode: l.head}
	l.size++
	if l.size == 1 {
		l.tail = l.head
	} else {
		l.head.NextNode.PreviousNode = l.head
	}
}

func (l *List) Push(v interface{}) {
	l.tail = &Node{Value: v, PreviousNode: l.tail}
	l.size++
	if l.size == 1 {
		l.head = l.tail
	} else {
		l.tail.PreviousNode.NextNode = l.tail
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.size == 0 {
		return nil, ErrEmptyList
	}
	value := l.head.Value
	if l.size == 1 {
		l.head, l.tail = nil, nil
	} else {
		l.head = l.head.NextNode
		l.head.PreviousNode = nil
	}
	l.size--
	return value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.size == 0 {
		return nil, ErrEmptyList
	}
	value := l.tail.Value
	if l.size == 1 {
		l.head, l.tail = nil, nil
	} else {
		l.tail = l.tail.PreviousNode
		l.tail.NextNode = nil
	}
	l.size--
	return value, nil
}

func (l *List) Reverse() {
	if l.size == 0 {
		return
	}
	l.tail, l.head = l.head, l.tail
	currentNode := l.head
	for currentNode != nil {
		currentNode.NextNode, currentNode.PreviousNode = currentNode.PreviousNode, currentNode.NextNode
		currentNode = currentNode.NextNode
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
