package alphametics

import (
	"errors"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	separateSides := strings.Split(puzzle, " == ")
	summands, sum := strings.Split(separateSides[0], " + "), separateSides[1]
	totalContributions := map[string]int{}
	isFirstDigit := map[string]bool{}
	for _, word := range summands {
		multiplicand := 1
		for i := len(word) - 1; i >= 0; i-- {
			letter := string(word[i])
			if i == 0 {
				isFirstDigit[letter] = true
			}
			totalContributions[letter] += multiplicand
			multiplicand *= 10
		}
	}
	multiplicand := 1
	for i := len(sum) - 1; i >= 0; i-- {
		letter := string(sum[i])
		if i == 0 {
			isFirstDigit[letter] = true
		}
		totalContributions[letter] -= multiplicand
		multiplicand *= 10
	}
	letters := []string{}
	for key := range totalContributions {
		letters = append(letters, key)
	}
	solutions := []map[string]int{}
	for i := 0; i < Factorial(10)/Factorial(10-len(letters)); i++ {
		letterValues := map[string]int{}
		iCopy := i
		firstDigitZero := false
		numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		sum := 0
		for index, letter := range letters {
			numbersIndex := iCopy % (10-index)
			numberToTry := numbers[numbersIndex]
			if numberToTry == 0 && isFirstDigit[letter] {
				firstDigitZero = true
				break
			}
			letterValues[letter] = numberToTry
			sum += totalContributions[letter] * letterValues[letter]
			numbers = append(numbers[:numbersIndex], numbers[numbersIndex+1:]...)
			iCopy /= (10-index)
		}
		if firstDigitZero {
			continue
		}
		if sum == 0 {
			solutions = append(solutions, letterValues)
		}
	}
	if len(solutions) != 1 {
		return nil, errors.New("no unique solution")
	} else {
		return solutions[0], nil
	}
}

func Factorial(input int) int {
	if input == 1 || input == 0 {
		return 1
	}
	return input * Factorial(input-1)
}
