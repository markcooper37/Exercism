package atbash

import (
	"fmt"
	"unicode"
)

func Atbash(s string) string {
	cipher := ""
	charsAdded := 0
	for _, char := range s {
		char = unicode.ToLower(char)
		if char > 47 && char < 58 {
			if charsAdded == 5 {
				cipher = fmt.Sprintf("%s%s", cipher, " ")
				charsAdded = 0
			}
			cipher = fmt.Sprintf("%s%s", cipher, string(char))
			charsAdded++
		} else if char > 96 && char < 123 {
			if charsAdded == 5 {
				cipher = fmt.Sprintf("%s%s", cipher, " ")
				charsAdded = 0
			}
			char = 219 - char
			cipher = fmt.Sprintf("%s%s", cipher, string(char))
			charsAdded++
		}
	}
	return cipher
}
