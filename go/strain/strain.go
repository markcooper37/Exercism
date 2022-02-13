package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var keepInts Ints
	for _, value := range i {
		if filter(value) {
			keepInts = append(keepInts, value)
		}
	}
	return keepInts
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var discardInts Ints
	for _, value := range i {
		if !filter(value) {
			discardInts = append(discardInts, value)
		}
	}
	return discardInts
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var keepLists Lists
	for _, value := range l {
		if filter(value) {
			keepLists = append(keepLists, value)
		}
	}
	return keepLists
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var keepStrings Strings
	for _, value := range s {
		if filter(value) {
			keepStrings = append(keepStrings, value)
		}
	}
	return keepStrings
}
