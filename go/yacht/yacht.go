package yacht

import "sort"

func Score(dice []int, category string) int {
	sort.Ints(dice)
	count := 0
	if category == "ones" {
		for _, die := range dice {
			if die == 1 {
				count++
			}
		}
	} else if category == "twos" {
		for _, die := range dice {
			if die == 2 {
				count += 2
			}
		}
	} else if category == "threes" {
		for _, die := range dice {
			if die == 3 {
				count += 3
			}
		}
	} else if category == "fours" {
		for _, die := range dice {
			if die == 4 {
				count += 4
			}
		}
	} else if category == "fives" {
		for _, die := range dice {
			if die == 5 {
				count += 5
			}
		}
	} else if category == "sixes" {
		for _, die := range dice {
			if die == 6 {
				count += 6
			}
		}
	} else if category == "full house" {
		if dice[0] == dice[1] && dice[3] == dice[4] && (dice[1] == dice[2] || dice[2] == dice[3]) && !(dice[1] == dice[2] && dice[2] == dice[3]) {
			for _, die := range dice {
				count += die
			}
		}
	} else if category == "four of a kind" {
		if dice[0] == dice[1] && dice[1] == dice[2] && dice[2] == dice[3] {
			count += dice[0] + dice[1] + dice[2] + dice[3]
		} else if dice[1] == dice[2] && dice[2] == dice[3] && dice[3] == dice[4] {
			count += + dice[1] + dice[2] + dice[3] + dice[4]
		}
	} else if category == "little straight" {
		if dice[0] == 1 && dice[1] == 2 && dice[2] == 3 && dice[3] == 4 && dice[4] == 5 {
			count = 30
		}
	} else if category == "big straight" {
		if dice[0] == 2 && dice[1] == 3 && dice[2] == 4 && dice[3] == 5 && dice[4] == 6 {
			count = 30
		}
	} else if category == "choice" {
		for _, die := range dice {
			count += die
		}
	} else if category == "yacht" {
		if dice[0] == dice[1] && dice[1] == dice[2] && dice[2] == dice[3] && dice[3] == dice[4] {
			count = 50
		}
	}
	return count
}
