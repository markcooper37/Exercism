package protein

import "errors"

var ErrStop = errors.New("found stop codon")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(rna)/3; i++ {
		codon := rna[3*i:3*i+3]
		protein, err := FromCodon(codon)
		if err == ErrInvalidBase {
			return nil, err
		}
		if err == ErrStop {
			return proteins, nil
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	switch {
	case codon == "AUG":
		return "Methionine", nil
	case codon == "UUU" || codon == "UUC":
		return "Phenylalanine", nil
	case codon == "UUA" || codon == "UUG":
		return "Leucine", nil
	case codon == "UCU" || codon == "UCC" || codon == "UCA" || codon == "UCG":
		return "Serine", nil
	case codon == "UAU" || codon == "UAC":
		return "Tyrosine", nil
	case codon == "UGU" || codon == "UGC":
		return "Cysteine", nil
	case codon == "UGG":
		return "Tryptophan", nil
	case codon == "UAA" || codon == "UAG" || codon == "UGA":
		return "", ErrStop
	}
	return "", ErrInvalidBase
}
