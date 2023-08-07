package linkedlist

import (
	"errors"
)

type Node struct {
	Value    interface{}
	next     *Node
	previous *Node
}

type List struct {
	first *Node
	last  *Node
	size  int
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
	return n.next
}

func (n *Node) Prev() *Node {
	return n.previous
}

func (l *List) Unshift(v interface{}) {
	l.first = &Node{Value: v, next: l.first}
	l.size++
	if l.size == 1 {
		l.last = l.first
	} else {
		l.first.next.previous = l.first
	}
}

func (l *List) Push(v interface{}) {
	l.last = &Node{Value: v, previous: l.last}
	l.size++
	if l.size == 1 {
		l.first = l.last
	} else {
		l.last.previous.next = l.last
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.size == 0 {
		return nil, ErrEmptyList
	}
	value := l.first.Value
	if l.size == 1 {
		l.first, l.last = nil, nil
	} else {
		l.first = l.first.next
		l.first.previous = nil
	}
	l.size--
	return value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.size == 0 {
		return nil, ErrEmptyList
	}
	value := l.last.Value
	if l.size == 1 {
		l.first, l.last = nil, nil
	} else {
		l.last = l.last.previous
		l.last.next = nil
	}
	l.size--
	return value, nil
}

func (l *List) Reverse() {
	if l.size == 0 {
		return
	}
	l.last, l.first = l.first, l.last
	current := l.first
	for current != nil {
		current.next, current.previous = current.previous, current.next
		current = current.next
	}
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
