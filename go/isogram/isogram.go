package isogram

import "strings"

func IsIsogram(word string) bool {
    wordLower := strings.ToLower(word)
	alphabet := map[string]bool{"a": false, "b": false, "c": false, "d": false,
                                "e": false, "f": false, "g": false, "h": false,
                                "i": false, "j": false, "k": false, "l": false,
                                "m": false, "n": false, "o": false, "p": false,
                                "q": false, "r": false, "s": false, "t": false,
                                "u": false, "v": false, "w": false, "x": false,
                                "y": false, "z": false}
    for _, char := range wordLower {
        charValue, charExists := alphabet[string(char)]
        if charExists {
            if charValue {
                return false
            } 
            alphabet[string(char)] = true
        }
    }
	return true
}
