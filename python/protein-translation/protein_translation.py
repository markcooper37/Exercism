PROTEIN_MAP = {
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan"
}


def proteins(strand):
    all_proteins = []
    for i in range(0, len(strand), 3):
        if strand[i:i + 3] in ["UAA", "UAG", "UGA"]:
            break
        all_proteins.append(PROTEIN_MAP[strand[i:i + 3]])
    return all_proteins
        
