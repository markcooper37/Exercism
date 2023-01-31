def to_rna(dna_strand):
    replacements = {"G": "C", "C": "G", "T": "A", "A": "U"}
    return "".join([replacements[letter] for letter in dna_strand])
