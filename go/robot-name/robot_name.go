package robotname

import (
	"errors"
	"fmt"
)

type Robot struct {
	name  int
	named bool
}

const allNames = 26 * 26 * 1000

var names = map[int]bool{}
var namesAssigned = 0

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (r *Robot) Name() (string, error) {
	if !r.named {
		if namesAssigned == allNames {
			return "", errors.New("all robot names have been assigned")
		}
		r.Reset()
	}
	return fmt.Sprintf("%c%c%03d", letters[r.name/26000], letters[(r.name%26000)/1000], r.name%1000), nil
}

func (r *Robot) Reset() {
	if len(names) == 0 {
		for i := 0; i < allNames; i++ {
			names[i] = false
		}
	}
	for name := range names {
		r.name = name
		break
	}
	delete(names, r.name)
	r.named = true
	namesAssigned++
}
