package bottlesong

import "fmt"

const (
	firstTwoLinesPattern = "%s green bottle%s hanging on the wall,"
	thirdLine            = "And if one green bottle should accidentally fall,"
	fourthLinePattern    = "There'll be %s green bottle%s hanging on the wall."
)

var firstBottleNumbers = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
var secondBottleNumbers = []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Recite(startBottles, takeDown int) []string {
	verses := []string{}
	for i := 0; i < takeDown; i++ {
		verses = append(verses, Verse(startBottles-i)...)
		if takeDown > 1 && i < takeDown-1 {
			verses = append(verses, "")
		}
	}
	return verses
}

func Verse(startBottles int) []string {
	verse := []string{}
	plural := "s"
	if startBottles == 1 {
		plural = ""
	}
	line := fmt.Sprintf(firstTwoLinesPattern, firstBottleNumbers[startBottles-1], plural)
	verse = append(verse, []string{line, line, thirdLine}...)
	if startBottles == 2 {
		plural = ""
	} else {
		plural = "s"
	}
	verse = append(verse, fmt.Sprintf(fourthLinePattern, secondBottleNumbers[startBottles-1], plural))
	return verse
}
