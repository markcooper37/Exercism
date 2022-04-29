package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

func (bst *BinarySearchTree) Insert(i int) {
	var newTree **BinarySearchTree
	if i > bst.data {
		newTree = &bst.right
	} else {
		newTree = &bst.left
	}
	if *newTree == nil {
		*newTree = NewBst(i)
	} else {
		(*newTree).Insert(i)
	}
}

func (bst *BinarySearchTree) SortedData() []int {
	result := []int{}
	if bst == nil {
		return result
	}
	result = append(result, bst.left.SortedData()...)
	result = append(result, bst.data)
	result = append(result, bst.right.SortedData()...)
	return result
}
