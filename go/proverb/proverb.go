// Package proverb contains a function that returns a proverb given a list of strings as input.
package proverb

import "fmt"

// Proverb takes a list of strings and returns a proverb based on these.
func Proverb(rhyme []string) []string {
	output := []string{}
    if len(rhyme) == 0 {
        return output
    }
	for i := 0; i < len(rhyme) - 1; i++ {
        output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
    }
	output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
    return output
}
