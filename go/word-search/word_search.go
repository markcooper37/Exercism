package wordsearch

import (
	"errors"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	solution := map[string][2][2]int{}
	for _, word := range words {
		wordFound := false
		for rowIndex, row := range puzzle {
			for columnIndex := range row {
				if word[0] == puzzle[rowIndex][columnIndex] {
					coordinates, err := CheckUp(word, puzzle, [2]int{columnIndex, rowIndex})
					if err == nil {
						solution[word] = coordinates
						wordFound = true
						break
					}
					coordinates, err = CheckDown(word, puzzle, [2]int{columnIndex, rowIndex})
					if err == nil {
						solution[word] = coordinates
						wordFound = true
						break
					}
					coordinates, err = CheckLeft(word, puzzle, [2]int{columnIndex, rowIndex})
					if err == nil {
						solution[word] = coordinates
						wordFound = true
						break
					}
					coordinates, err = CheckRight(word, puzzle, [2]int{columnIndex, rowIndex})
					if err == nil {
						solution[word] = coordinates
						wordFound = true
						break
					}
					coordinates, err = CheckUpLeft(word, puzzle, [2]int{columnIndex, rowIndex})
					if err == nil {
						solution[word] = coordinates
						wordFound = true
						break
					}
					coordinates, err = CheckUpRight(word, puzzle, [2]int{columnIndex, rowIndex})
					if err == nil {
						solution[word] = coordinates
						wordFound = true
						break
					}
					coordinates, err = CheckDownLeft(word, puzzle, [2]int{columnIndex, rowIndex})
					if err == nil {
						solution[word] = coordinates
						wordFound = true
						break
					}
					coordinates, err = CheckDownRight(word, puzzle, [2]int{columnIndex, rowIndex})
					if err == nil {
						solution[word] = coordinates
						wordFound = true
						break
					}
				}
			}
			if wordFound {
				break
			}
		}
		if !wordFound {
			return nil, errors.New("word is not contained in the puzzle")
		}
	}
	return solution, nil
}

func CheckUp(word string, puzzle []string, index [2]int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[1] < 0 {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			currentIndex[1]--
		}
	}
	return [2][2]int{index, currentIndex}, nil
}

func CheckDown(word string, puzzle []string, index [2]int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[1] >= len(puzzle) {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			currentIndex[1]++
		}
	}
	return [2][2]int{index, currentIndex}, nil
}

func CheckLeft(word string, puzzle []string, index [2]int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[0] < 0 {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			currentIndex[0]--
		}
	}
	return [2][2]int{index, currentIndex}, nil
}

func CheckRight(word string, puzzle []string, index [2]int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[0] >= len(puzzle[index[1]]) {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			currentIndex[0]++
		}
	}
	return [2][2]int{index, currentIndex}, nil
}

func CheckUpLeft(word string, puzzle []string, index [2]int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[1] < 0 || currentIndex[0] < 0 {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			currentIndex[1]--
			currentIndex[0]--
		}
	}
	return [2][2]int{index, currentIndex}, nil
}

func CheckUpRight(word string, puzzle []string, index [2]int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[1] < 0 || currentIndex[0] >= len(puzzle[index[1]]) {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			currentIndex[1]--
			currentIndex[0]++
		}
	}
	return [2][2]int{index, currentIndex}, nil
}

func CheckDownLeft(word string, puzzle []string, index [2]int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[1] >= len(puzzle) || currentIndex[0] < 0 {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			currentIndex[1]++
			currentIndex[0]--
		}
	}
	return [2][2]int{index, currentIndex}, nil
}

func CheckDownRight(word string, puzzle []string, index [2]int) ([2][2]int, error) {
	currentIndex := index
	for letterIndex := range word {
		if currentIndex[1] >= len(puzzle) || currentIndex[0] >= len(puzzle[index[1]]) {
			return [2][2]int{}, errors.New("index out of range")
		}
		if word[letterIndex] != puzzle[currentIndex[1]][currentIndex[0]] {
			return [2][2]int{}, errors.New("word not found")
		}
		if letterIndex != len(word)-1 {
			currentIndex[1]++
			currentIndex[0]++
		}
	}
	return [2][2]int{index, currentIndex}, nil
}
