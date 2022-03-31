package house

import (
	"fmt"
	"strings"
)

var phrases = [][]string{{"malt", "lay in"}, {"rat", "ate"}, {"cat", "killed"}, {"dog", "worried"}, {"cow with the crumpled horn", "tossed"}, {"maiden all forlorn", "milked"}, {"man all tattered and torn", "kissed"}, {"priest all shaven and shorn", "married"}, {"rooster that crowed in the morn", "woke"}, {"farmer sowing his corn", "kept"}, {"horse and the hound and the horn", "belonged to"}}

func Verse(v int) string {
	if v < 1 || v > 12 {
		return ""
	} else if v == 1 {
		return "This is the house that Jack built."
	} else {
		adjustedPreviousVerse := strings.Replace(Verse(v-1), "This is", fmt.Sprintf("that %s", phrases[v-2][1]), 1)
		return fmt.Sprintf("This is the %s\n%s", phrases[v-2][0], adjustedPreviousVerse)
	}
}

func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		if i > 1 {
			song = fmt.Sprintf("%s\n\n", song)
		}
		song = fmt.Sprintf("%s%s", song, Verse(i))
	}
	return song
}
