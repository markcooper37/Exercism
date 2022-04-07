package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("invalid input character")
	}
	lineCount := 2*int(char-'A') + 1
	lines := make([]string, lineCount)
	line := strings.Repeat(" ", int(char-'A')) + "A" + strings.Repeat(" ", int(char-'A')) + "\n"
	lines[0] = line
	lines[lineCount-1] = line
	for i := 1; i <= int(char-'A'); i++ {
		line := strings.Repeat(" ", int(char-'A')-i) + string('A'+byte(i)) + strings.Repeat(" ", 2*i-1) + string('A'+byte(i)) + strings.Repeat(" ", int(char-'A')-i) + "\n"
		lines[i] = line
		lines[lineCount-i-1] = line

	}
	diamond := ""
	for i := 0; i < lineCount; i++ {
		diamond += lines[i]
	}
	return diamond, nil
}
