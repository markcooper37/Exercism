package thefarm

import (
	"errors"
	"fmt"
)

// SillyNephewError implements the error interface to return an error
// message when the number of cows is negative
type SillyNephewError struct {
	cows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if err != nil {
		if err != ErrScaleMalfunction {
			return 0, err
		}
		if fodder > 0 {
			fodder *= 2
		}
	}
	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		sillyNephewError := SillyNephewError{cows: cows}
		return 0, &sillyNephewError
	}
	return fodder / float64(cows), nil
}
