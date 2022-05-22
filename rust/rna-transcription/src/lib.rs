use std::collections::HashMap;

#[derive(Debug, PartialEq)]
pub struct Dna {
    dna: String
}

#[derive(Debug, PartialEq)]
pub struct Rna {
    rna: String
}

impl Dna {
    pub fn new(dna: &str) -> Result<Dna, usize> {
        for char in dna.chars().enumerate() {
            if char.1 != 'G' && char.1 != 'C' && char.1 != 'T' && char.1 != 'A' {
                return Err(char.0)
            }
        }
        Ok(Dna{dna: dna.to_string()})
    }

    pub fn into_rna(self) -> Rna {
        let mut complements = HashMap::new();
        complements.insert('G', 'C');
        complements.insert('C', 'G');
        complements.insert('T', 'A');
        complements.insert('A', 'U');
        Rna{rna: self.dna.chars().map(|c| complements.get(&c).unwrap()).collect::<String>()}
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        for char in rna.chars().enumerate() {
            if char.1 != 'G' && char.1 != 'C' && char.1 != 'U' && char.1 != 'A' {
                return Err(char.0)
            }
        }
        Ok(Rna{rna: rna.to_string()})
    }
}
