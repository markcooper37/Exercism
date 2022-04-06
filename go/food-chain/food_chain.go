package foodchain

import "fmt"

var lines = []string{"It wriggled and jiggled and tickled inside her.\n", "How absurd to swallow a bird!\n", "Imagine that, to swallow a cat!\n", "What a hog, to swallow a dog!\n", "Just opened her throat and swallowed a goat!\n", "I don't know how she swallowed a cow!\n"}
var animals = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow"}

func Verse(v int) string {
	if v <= 0 || v > 8 {
		return ""
	} else if v == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	} else {
		verse := fmt.Sprintf("I know an old lady who swallowed a %s.\n", animals[v-1])
		if v > 1 {
			verse += lines[v-2]
		}
		for i := v-2; i >= 0; i-- {
			if i == 1 {
				verse += "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.\n"
			} else {
				verse += fmt.Sprintf("She swallowed the %s to catch the %s.\n", animals[i+1], animals[i])
			}
		}
		return verse + "I don't know why she swallowed the fly. Perhaps she'll die."
	}
}

func Verses(start, end int) string {
	if start > end || start < 1 || end > 8 {
		return ""
	} else {
		verses := Verse(start)
		for i := start + 1; i <= end; i++ {
			verses += "\n\n" + Verse(i)
		}
		return verses
	}
}

func Song() string {
	return Verses(1, 8)
}
