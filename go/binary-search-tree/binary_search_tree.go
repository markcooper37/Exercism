package binarysearchtree

import "strconv"

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
	return SearchTreeData{data: i}
}

// Insert inserts an int into the SearchTreeData.
func (std *SearchTreeData) Insert(i int) {
	if std == nil {
		*std = NewBst(i)
		return
	}
	currentNode := std
	for {
		if i <= currentNode.data {
			if currentNode.left == nil {
				currentNode.left = &SearchTreeData{data: i}
				return
			} else {
				currentNode = currentNode.left
			}
		} else {
			if currentNode.right == nil {
				currentNode.right = &SearchTreeData{data: i}
				return
			} else {
				currentNode = currentNode.right
			}
		}
	}
}

// MapString returns the ordered contents of SearchTreeData as a []string.
func (std *SearchTreeData) MapString(func(int) string) (result []string) {
	if std == nil {
		return []string{}
	}
	return std.StringArrayFromChildren()
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
func (std *SearchTreeData) MapInt(func(int) int) (result []int) {
	if std == nil {
		return []int{}
	}
	return std.IntArrayFromChildren()
}

func (std *SearchTreeData) IntArrayFromChildren() []int {
	if std.left == nil && std.right == nil {
		return []int{std.data}
	} else if std.left == nil {
		return append([]int{std.data}, std.right.IntArrayFromChildren()...)
	} else if std.right == nil {
		return append(std.left.IntArrayFromChildren(), std.data)
	}
	array := append(std.left.IntArrayFromChildren(), std.data)
	array = append(array, std.right.IntArrayFromChildren()...)
	return array
}

func (std *SearchTreeData) StringArrayFromChildren() []string {
	if std.left == nil && std.right == nil {
		return []string{strconv.Itoa(std.data)}
	} else if std.left == nil {
		return append([]string{strconv.Itoa(std.data)}, std.right.StringArrayFromChildren()...)
	} else if std.right == nil {
		return append(std.left.StringArrayFromChildren(), strconv.Itoa(std.data))
	}
	array := append(std.left.StringArrayFromChildren(), strconv.Itoa(std.data))
	array = append(array, std.right.StringArrayFromChildren()...)
	return array
}
