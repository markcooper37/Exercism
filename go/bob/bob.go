package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	containsLetters := false
	for _, char := range remark {
		if unicode.IsLetter(char) {
			containsLetters = true
		}
	}
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	} else if string(remark[len(remark)-1]) == "?" && remark == strings.ToUpper(remark) && containsLetters {
		return "Calm down, I know what I'm doing!"
	} else if string(remark[len(remark)-1]) == "?" {
		return "Sure."
	} else if remark == strings.ToUpper(remark) && containsLetters {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
