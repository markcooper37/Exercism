package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
    h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
    for _, char := range d {
        if value, ok := h[char]; ok {
        	h[char] = value + 1
        } else {
        	var emptyHist Histogram
        	return emptyHist, errors.New("input contains invalid character")
        }
    }
	return h, nil
}