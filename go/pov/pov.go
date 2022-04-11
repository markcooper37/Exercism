package pov

type Tree struct {
	value    string
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{value: value, children: children}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	oldRootToNewRoot := tr.FindTrees(from)
	if oldRootToNewRoot == nil {
		return nil
	}
	for i := 0; i < len(oldRootToNewRoot) - 1; i++ {
		children := []*Tree{}
		for _, child := range oldRootToNewRoot[i].children {
			if child != oldRootToNewRoot[i+1] {
				children = append(children, child)
			}
		}
		oldRootToNewRoot[i].children = children
		oldRootToNewRoot[i+1].children = append(oldRootToNewRoot[i+1].children, oldRootToNewRoot[i])
	}
	return oldRootToNewRoot[len(oldRootToNewRoot)-1]
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	newTree := tr.FromPov(from)
	if newTree == nil {
		return nil
	}
	path := []string{from}
	for _, child := range newTree.children {
		if child.value == to {
			path = append(path, to)
			return path
		} else {
			remainingPath := child.PathTo(child.value, to)
			if remainingPath != nil {
				path = append(path, remainingPath...)
				return path
			}
		}
	}
	return nil
}

func (tr *Tree) FindTrees(value string) []*Tree {
	if tr.value == value {
		return []*Tree{tr}
	}
	for _, child := range tr.children {
		requiredTrees := child.FindTrees(value)
		if requiredTrees != nil {
			return append([]*Tree{tr}, requiredTrees...)
		}
	}
	return nil
}
