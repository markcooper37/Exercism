package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	output := ""
	count := 1
	for index, char := range input {
		if index < len(input)-1 && char == rune(input[index+1]) {
			count++
			continue
		} else {
			if count > 1 {
				output += strconv.Itoa(count)
			}
			output += string(char)
			count = 1
		}
	}
	return output
}

func RunLengthDecode(input string) string {
	output := ""
	length := 0
	for _, char := range input {
		if value, err := strconv.Atoi(string(char)); err == nil {
			length += 9 * length + value 
		} else {
			if length == 0 {
				output += string(char)
			} else {
				output += strings.Repeat(string(char), length)
			}
			length = 0
		}
	}
	return output
}
