use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    if nucleotide != 'A' && nucleotide != 'C' && nucleotide != 'G' && nucleotide != 'T' {
        return Err(nucleotide);
    }
    let mut count = 0;
    for character in dna.chars() {
        if character != 'A' && character != 'C' && character != 'G' && character != 'T' {
            return Err(character);
        } else if character == nucleotide {
            count += 1
        }
    }
    Ok(count)
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let nucleotides = ['A', 'T', 'C', 'G'];
    let mut counts = HashMap::new();
    for nucleotide in nucleotides {
        match count(nucleotide, dna) {
            Ok(count) => {counts.insert(nucleotide, count);}
            Err(character) => return Err(character)
        }
    }
    Ok(counts)
}
