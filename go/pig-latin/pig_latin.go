package piglatin

import "strings"

func Sentence(sentence string) string {
	splitSentence := strings.Split(sentence, " ")
	words := []string{}
	for _, word := range splitSentence {
		pos := 0
		for pos = 0; pos < len(word); pos++ {
			if pos == 0 && len(word) > 2 && (word[0:2] == "xr" || word[0:2] == "yt") {
				break
			} else if word[pos] == 'q' && pos < len(word)-1 && word[pos+1] == 'u' {
				pos += 2
			} else if pos > 0 && word[pos] == 'y' {
				break
			} else if word[pos] != 'a' && word[pos] != 'e' && word[pos] != 'i' && word[pos] != 'o' && word[pos] != 'u' {
				continue
			}
			break
		}
		words = append(words, word[pos:]+word[:pos]+"ay")
	}
	return strings.Join(words, " ")
}
