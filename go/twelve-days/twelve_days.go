package twelve

import "fmt"

var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", 
                    "ninth", "tenth", "eleventh", "twelfth"}
var gifts = []string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", 
                     "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", 
                     "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", 
                     "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

func Verse(i int) string {
    verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i - 1])
    for j := i - 1; j > 0; j-- {
        verse = fmt.Sprintf("%s %s,", verse, gifts[j])
    }
	if i > 1 {
        verse = fmt.Sprintf("%s and", verse)
    }
	verse = fmt.Sprintf("%s %s.", verse, gifts[0])
    return verse
}

func Song() string {
	song := ""
    for i := 1; i < 12; i++ {
        song = fmt.Sprintf("%s%s\n", song, Verse(i))
    }
	song = fmt.Sprintf("%s%s", song, Verse(12))
	return song
}
