package stringset

import "fmt"

type Set map[string]bool

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, sliceElement := range l {
		for setElement := range s {
			if setElement == sliceElement {
				break
			}
		}
		s.Add(sliceElement)
	}
	return s
}

func (s Set) String() string {
	setString := "{"
	addedAnElement := false
	for setElement := range s {
		if addedAnElement {
			setString = fmt.Sprintf("%s, ", setString)
		}
		addedAnElement = true
		setString = fmt.Sprintf("%s\"%s\"", setString, setElement)
	}
	setString = fmt.Sprintf("%s}", setString)
	return setString
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	for setElement := range s {
		if elem == setElement {
			return true
		}
	}
	return false
}

func (s Set) Add(elem string) {
	for setElement := range s {
		if elem == setElement {
			return
		}
	}
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for s1Element := range s1 {
		setsEqual := false
		for s2Element := range s2 {
			if s1Element == s2Element {
				setsEqual = true
				break
			}
		}
		if !setsEqual {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for s1Element := range s1 {
		for s2Element := range s2 {
			if s1Element == s2Element {
				return false
			}
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	s := New()
	for s1Element := range s1 {
		for s2Element := range s2 {
			if s1Element == s2Element {
				s.Add(s1Element)
				break
			}
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for s1Element := range s1 {
		inS2 := false
		for s2Element := range s2 {
			if s1Element == s2Element {
				inS2 = true
				break
			}
		}
		if !inS2 {
			s.Add(s1Element)
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	s := New()
	for s1Element := range s1 {
		s.Add(s1Element)
	}
	for s2Element := range s2 {
		s.Add(s2Element)
	}
	return s
}
