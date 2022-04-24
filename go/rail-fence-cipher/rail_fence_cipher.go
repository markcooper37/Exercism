package railfence

import "strings"

func Encode(message string, rails int) string {
	pieces := make([]string, rails)
	line := 0
	step := 1
	for _, character := range message {
		if character == ' ' {
			continue
		}
		pieces[line] += string(character)
		line += step
		if line == 0 || line == rails-1 {
			step *= -1
		}
	}
	return strings.Join(pieces, "")
}

func Decode(message string, rails int) string {
	lengths := make([]int, rails)
	line := 0
	step := 1
	for i := 0; i < len(message); i++ {
		lengths[line]++
		line += step
		if line == 0 || line == rails-1 {
			step *= -1
		}
	}
	pieces := make([]string, rails)
	start := 0
	for index, length := range lengths {
		pieces[index] = message[start:start+length]
		start += length
	}
	indices := make([]int, rails)
	decoded := ""
	line = 0
	step = 1
	for i := 0; i < len(message); i++ {
		decoded += string(pieces[line][indices[line]])
		indices[line]++
		line += step
		if line == 0 || line == rails-1 {
			step *= -1
		}
	}
	return decoded
}
