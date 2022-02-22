package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	question = strings.Trim(question, "?")
	splitQuestion := strings.Split(question, " ")
	if len(splitQuestion) < 3 {
		return 0, false
	}
	if splitQuestion[0] != "What" || splitQuestion[1] != "is" {
		return 0, false
	}
	initialValue, err := strconv.Atoi(splitQuestion[2])
	if err != nil {
		return 0, false
	}
	for i := 3; i < len(splitQuestion); {
		switch splitQuestion[i] {
		case "plus":
			if len(splitQuestion) < i+2 {
				return 0, false
			}
			operand, err := strconv.Atoi(splitQuestion[i+1])
			if err != nil {
				return 0, false
			}
			initialValue += operand
			i += 2
		case "minus":
			if len(splitQuestion) < i+2 {
				return 0, false
			}
			operand, err := strconv.Atoi(splitQuestion[i+1])
			if err != nil {
				return 0, false
			}
			initialValue -= operand
			i += 2
		case "multiplied":
			if len(splitQuestion) < i+3 || splitQuestion[i+1] != "by" {
				return 0, false
			}
			operand, err := strconv.Atoi(splitQuestion[i+2])
			if err != nil {
				return 0, false
			}
			initialValue *= operand
			i += 3
		case "divided":
			if len(splitQuestion) < i+3 || splitQuestion[i+1] != "by" {
				return 0, false
			}
			operand, err := strconv.Atoi(splitQuestion[i+2])
			if err != nil {
				return 0, false
			}
			initialValue /= operand
			i += 3
		default:
			return 0, false
		}
	}
	return initialValue, true
}
