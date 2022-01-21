package luhn

import (
    "strings"
    "unicode"
    "strconv"
)

func Valid(id string) bool {
	newId := strings.ReplaceAll(id, " ", "")
    if len(newId) < 2 {
        return false
    }
	digits := make([]int, len(newId))
    for i, char := range newId {
        if !unicode.IsDigit(char) {
            return false
        }
    	charAsInt, _ := strconv.Atoi(string(char))
    	digits[i] = charAsInt
    }
	for j := len(newId) - 2; j >= 0; j = j - 2 {
        newDigit := digits[j] * 2
        if newDigit > 9 {
            newDigit = newDigit - 9
        }
    	digits[j] = newDigit
    }
	sum := 0
	for k := range digits {
        sum = sum + digits[k]
    }
	return sum % 10 == 0
}
