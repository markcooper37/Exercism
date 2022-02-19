package scale

import "unicode"

var majorScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var minorScale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var steps = map[rune]int{'m': 1, 'M': 2, 'A': 3}

func Scale(tonic, interval string) []string {
	var scale []string
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		tonic = firstLetterToUpper(tonic)
		position := findFirstPosition(tonic, majorScale)
		scale = findScale(position, interval, majorScale)
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		tonic = firstLetterToUpper(tonic)
		position := findFirstPosition(tonic, minorScale)
		scale = findScale(position, interval, minorScale)
	}
	return scale
}

func findFirstPosition(inputTonic string, scale []string) int {
	for index, tonic := range scale {
		if inputTonic == tonic {
			return index
		}
	}
	return -1
}

func firstLetterToUpper(input string) string {
	r := []rune(input)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func findScale(position int, interval string, scale []string) []string {
	var finalScale []string
	if interval == "" {
		for i := 0; i < len(scale); i++ {
			finalScale = append(finalScale, scale[position])
			position = (position + 1) % len(scale)
		}
	} else {
		for _, char := range interval {
			finalScale = append(finalScale, scale[position])
			position = (position + steps[char]) % len(scale)
		}
	}
	return finalScale
}
