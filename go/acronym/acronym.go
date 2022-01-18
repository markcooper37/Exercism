// Package acronym implements a function to convert a phrase to its acronym.
package acronym

import (
    "unicode"
    "strings"
    "fmt"
)

// Abbreviate converts a phrase to its acronym
func Abbreviate(s string) string {
	f := func(c rune) bool {
        return !unicode.IsLetter(c) && string(c) != "'"
    }
	acronym := ""
    words := strings.FieldsFunc(s, f)
    for _, word := range words {
        acronym = fmt.Sprintf("%s%s", acronym, strings.ToUpper(string(word[0])))
    }
	return acronym
}
