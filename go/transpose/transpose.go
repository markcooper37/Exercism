package transpose

import "strings"

func Transpose(input []string) []string {
	output := []string{}
	for lineNumber, line := range input {
		for index, character := range line {
			if index == len(output) {
				output = append(output, strings.Repeat(" ", lineNumber))
			}
			if len(output[index]) < lineNumber {
				output[index] += strings.Repeat(" ", lineNumber - len(output[index]))
			}
			output[index] += string(character)
		}
	}
	return output
}
