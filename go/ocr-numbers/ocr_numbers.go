package ocr

import (
	"strings"
)

func Recognize(input string) []string {
	separateLines := strings.Split(input, "\n")
	numbers := []string{}
	for i := 1; i < len(separateLines); i += 4 {
		number := ""
		for j := 0; j < len(separateLines[i]); j += 3 {
			digit := separateLines[i][j:j+3] + "\n" + separateLines[i+1][j:j+3] + "\n" + separateLines[i+2][j:j+3]
			number += recognizeDigit(digit)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func recognizeDigit(digit string) string {
	switch digit {
	case " _ \n| |\n|_|":
		return "0"
	case "   \n  |\n  |":
		return "1"
	case " _ \n _|\n|_ ":
		return "2"
	case " _ \n _|\n _|":
		return "3"
	case "   \n|_|\n  |":
		return "4"
	case " _ \n|_ \n _|":
		return "5"
	case " _ \n|_ \n|_|":
		return "6"
	case " _ \n  |\n  |":
		return "7"
	case " _ \n|_|\n|_|":
		return "8"
	case " _ \n|_|\n _|":
		return "9"
	default:
		return "?"
	}
}
