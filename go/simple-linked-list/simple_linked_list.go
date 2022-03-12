package linkedlist

import (
	"errors"
)

// Define the List and Element types here.
type Element struct {
	value int
	next  *Element
}
type List []Element

func New(input []int) *List {
	newList := List{}
	if len(input) > 0 {
		newList = append(newList, Element{value: input[0]})
	}
	for i := 1; i < len(input); i++ {
		newList = append(newList, Element{value: input[i]})
		newList[i-1].next = &newList[i]
	}
	return &newList
}

func (l *List) Size() int {
	return len(*l)
}

func (l *List) Push(element int) {
	*l = append(*l, Element{value: element})
	if len(*l) > 1 {
		(*l)[len(*l)-2].next = &(*l)[len(*l)-1]
	}

}

func (l *List) Pop() (int, error) {
	if len(*l) == 0 {
		return 0, errors.New("list contains no elements")
	}
	lastValue := (*l)[len(*l)-1].value
	*l = (*l)[:len(*l)-1]
	if len(*l) > 0 {
		(*l)[len(*l)-1].next = nil
	}
	return lastValue, nil
}

func (l *List) Array() []int {
	array := []int{}
	for _, element := range *l {
		array = append(array, element.value)
	}
	return array
}

func (l *List) Reverse() *List {
	array := l.Array()
	reversedArray := []int{}
	for i := len(array) - 1; i >= 0; i-- {
		reversedArray = append(reversedArray, array[i])
	}
	return New(reversedArray)
}
