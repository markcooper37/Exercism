package railfence

import "strings"

func Encode(message string, rails int) string {
	pieces := make([]string, rails)
	line := 0
	inc := true
	for _, character := range message {
		if character == ' ' {
			continue
		}
		pieces[line] += string(character)
		if inc {
			line++
		} else {
			line--
		}
		if line == rails-1 {
			inc = false
		} else if line == 0 {
			inc = true
		}
	}
	return strings.Join(pieces, "")
}

func Decode(message string, rails int) string {
	return ""
}
