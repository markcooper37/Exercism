package alphametics

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	separateSides := strings.Split(puzzle, " == ")
	summands, sum := strings.Split(separateSides[0], " + "), separateSides[1]
	totalValues := map[string]int{}
	isFirstDigit := map[string]bool{}
	for _, word := range summands {
		multiplicand := 1
		for i := len(word) - 1; i >= 0; i-- {
			letter := string(word[i])
			if i == 0 && len(word) > 1 {
				isFirstDigit[letter] = true
			}
			totalValues[letter] += multiplicand
			multiplicand *= 10
		}
	}
	multiplicand := 1
	for i := len(sum) - 1; i >= 0; i-- {
		letter := string(sum[i])
		if i == 0 && len(sum) > 1 {
			isFirstDigit[letter] = true
		}
		totalValues[letter] -= multiplicand
		multiplicand *= 10
	}

	letters := []string{}
	for key := range totalValues {
		letters = append(letters, key)
	}
	solutions := []map[string]int{}
	for i := 0; i < int(math.Pow(10, float64(len(letters)))); i++ {
		fmt.Println(letters, i)
		numberUsed := map[int]bool{}
		letterValues := map[string]int{}
		numberToTest := i
		duplicateNumber := false
		firstDigitZero := false
		for _, letter := range letters {
			numberToTry := numberToTest % 10
			if numberUsed[numberToTry] {
				duplicateNumber = true
				break
			}
			if numberToTry == 0 && isFirstDigit[letter] {
				firstDigitZero = true
				break
			}
			numberUsed[numberToTry] = true
			letterValues[letter] = numberToTry
			numberToTest /= 10
		}
		if duplicateNumber || firstDigitZero {
			continue
		}
		sum := 0
		for _, letter := range letters {
			sum += totalValues[letter] * letterValues[letter]
		}
		if sum == 0 {
			solutions = append(solutions, letterValues)
		}
	}
	if len(solutions) != 1 {
		fmt.Println(solutions)
		return nil, errors.New("no unique solution")
	} else {
		return solutions[0], nil
	}
}
