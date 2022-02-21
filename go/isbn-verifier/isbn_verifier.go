package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	factor := 10
	sum := 0
	for i := 0; i < len(isbn)-1; i++ {
		integerValue, err := strconv.Atoi(string(isbn[i]))
		if err != nil {
			return false
		}
		sum += factor * integerValue
		factor--
	}
	if isbn[len(isbn)-1] == 'X' {
		sum += 10
	} else {
		integerValue, err := strconv.Atoi(string(isbn[len(isbn)-1]))
		if err != nil {
			return false
		}
		sum += integerValue
	}
	return sum%11 == 0
}
