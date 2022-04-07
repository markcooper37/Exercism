package change

import (
	"errors"
)

func Change(coins []int, target int) ([]int, error) {
	change := []int{}
	if target < 0 {
		return change, errors.New("target cannot be negative")
	} else if target == 0 {
		return change, nil
	}
	for i := len(coins) - 1; i >= 0; i-- {
		best, err := Change(coins[:i+1], target - coins[i])
		if err == nil && (len(best) < len(change)-1 || len(change) == 0) {
			change = append(best, coins[i])
		}
		if len(change) > 0 && i != 0 && target/coins[i-1] >= len(change) {
			break
		}
	}
	if len(change) == 0 {
		return change, errors.New("exact change cannot be given")
	}
	return change, nil
}
