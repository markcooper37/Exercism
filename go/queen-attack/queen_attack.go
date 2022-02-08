package queenattack

import "errors"

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if !validPositions(whitePosition, blackPosition) {
		return false, errors.New("invalid input")
	}
	if whitePosition[0] == blackPosition[0] || whitePosition[1] == blackPosition[1] {
		return true, nil
	}
	if int(whitePosition[0]) - int(whitePosition[1]) == int(blackPosition[0]) - int(blackPosition[1]) {
		return true, nil
	}
	if int(whitePosition[0]) + int(whitePosition[1]) == int(blackPosition[0]) + int(blackPosition[1]) {
		return true, nil
	}
	return false, nil
}

func validPositions(whitePosition, blackPosition string) bool {
	if len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false
	}
	if whitePosition[0] == blackPosition[0] && whitePosition[1] == blackPosition[1] {
		return false
	}
	if int(whitePosition[0]) < 97 || int(whitePosition[0]) > 104 {
		return false
	}
	if int(blackPosition[0]) < 97 || int(blackPosition[0]) > 104 {
		return false
	}
	if int(whitePosition[1]) < 49 || int(whitePosition[1]) > 56 {
		return false
	}
	if int(blackPosition[1]) < 49 || int(blackPosition[1]) > 56 {
		return false
	}
	return true
}
