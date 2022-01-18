package beer

import (
    "fmt"
    "errors"
)

func Song() string {
	song, _ := Verses(99, 0)
    return song
}

func Verses(start, stop int) (string, error) {
    if start < stop {
        return "", errors.New("Start input must be at least as big as stop input")
    }
	if start < 0 || start > 99 || stop < 0 || stop > 99 {
        return "", errors.New("Start and stop inputs must be between 0 and 99")
    }
    songSection := ""
	for i := start; i >= stop; i-- {
        verse, _ := Verse(i)
        songSection = fmt.Sprintf("%s%s%s", songSection, verse, "\n")
    }
	return songSection, nil
}

func Verse(n int) (string, error) {
    switch {
    case n > 2 && n < 100:
    	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
    case n == 2:
    	return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
    case n == 1:
    	return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
    case n == 0:
    	return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
    default:
    	return "", errors.New("Verse input must be between 0 and 99")
    }
}
