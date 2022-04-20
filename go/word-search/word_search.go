package wordsearch

import (
	"errors"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	solution := map[string][2][2]int{}
	for _, word := range words {
		for rowIndex, row := range puzzle {
			for columnIndex := range row {
				for i := 0; i <= 7; i++ {
					coordinates, err := CheckOrientation(word, puzzle, [2]int{columnIndex, rowIndex}, i)
					if err == nil {
						solution[word] = coordinates
					}
				}
			}
		}
	}
	if len(solution) != len(words) {
		return nil, errors.New("word missing from puzzle")
	}
	return solution, nil
}

func CheckOrientation(word string, puzzle []string, index [2]int, orientation int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[1] < 0 || currentIndex[1] >= len(puzzle) || currentIndex[0] < 0 || currentIndex[0] >= len(puzzle[index[1]]) {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			if orientation == 0 || orientation == 1 || orientation == 2 {
				currentIndex[1]--
			}
			if orientation == 2 || orientation == 3 || orientation == 4 {
				currentIndex[0]++
			}
			if orientation == 4 || orientation == 5 || orientation == 6 {
				currentIndex[1]++
			}
			if orientation == 6 || orientation == 7 || orientation == 0 {
				currentIndex[0]--
			}
		}
	}
	return [2][2]int{index, currentIndex}, nil
}
