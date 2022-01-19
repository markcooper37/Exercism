package wordcount

import (
    "regexp"
    "strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    wordCount := Frequency{}
	re := regexp.MustCompile(`[a-zA-Z]+'[a-zA-Z]+|[0-9]+|[a-zA-Z]+`)
    words := re.FindAllString(phrase, -1)
    for _, word := range words {
        wordCount[strings.ToLower(word)] += 1
    }
	return wordCount
}
