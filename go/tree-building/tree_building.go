package tree

import (
	"errors"
)

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	root := Node{}
	for i := 0; i < len(records); i++ {
		recordFound := false
		for _, record := range records {
			if record.ID == i {
				recordFound = true
				newNode := Node{ID: record.ID}
				if record.ID == 0 {
					if record.Parent != 0 {
						return nil, errors.New("root node has parent")
					}
					root = newNode
				} else {
					parent := FindParent(record.Parent, &root)
					if parent == nil {
						return nil, errors.New("parent does not exist")
					}
					parent.Children = append(parent.Children, &newNode)
				}
				break
			}
		}
		if !recordFound {
			return nil, errors.New("record does not exist")
		}
	}
	return &root, nil
}

func FindParent(parentID int, node *Node) *Node {
	if node.ID == parentID {
		return node
	}
	for _, child := range node.Children {
		parent := FindParent(parentID, child)
    if parent != nil {
			return parent
		}
	}
	return nil
}
