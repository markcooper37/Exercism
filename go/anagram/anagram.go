package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	letterFound := map[int]bool{}
	var anagrams []string
	subject = strings.ToLower(subject)
	var newCandidates []string
	for i := range candidates {
		newCandidates = append(newCandidates, strings.ToLower(candidates[i]))
	}
	subjectCharacters := make([]rune, len(subject))
	for index, character := range subject {
		subjectCharacters[index] = character
	}
	for candidateIndex, candidate := range newCandidates {
		if len(candidate) != len(subject) || candidate == subject {
			continue
		}
		candidateValid := true
		for i := 0; i < len(subject); i++ {
			letterFound[i] = false
		}
		for _, candidateCharacter := range candidate {
			letterInSubject := false
			for subjectIndex, subjectCharater := range subject {
				if letterFound[subjectIndex] {
					continue
				}
				if candidateCharacter == subjectCharater {
					letterFound[subjectIndex] = true
					letterInSubject = true
					break
				}
			}
			if !letterInSubject {
				candidateValid = false
			}
		}
		if candidateValid {
			anagrams = append(anagrams, candidates[candidateIndex])
		}
	}
	return anagrams
}
