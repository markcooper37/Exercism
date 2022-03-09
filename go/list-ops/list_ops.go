package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, value := range s {
		initial = fn(initial, value)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for index := s.Length()-1; index >= 0; index-- {
		initial = fn(s[index], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	newList := IntList{}
	for _, value := range s {
		if fn(value) {
			newList = append(newList, value)
		}
	}
	return newList
}

func (s IntList) Length() int {
	count := 0
	for range s {
		count++
	}
	return count
}

func (s IntList) Map(fn func(int) int) IntList {
	newList := IntList{}
	for _, value := range s {
		newList = append(newList, fn(value))
	}
	return newList
}

func (s IntList) Reverse() IntList {
	newList := IntList{}
	for index := s.Length() - 1; index >= 0; index-- {
		newList = append(newList, s[index])
	}
	return newList
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = s.Append(list)
	}
	return s
}
