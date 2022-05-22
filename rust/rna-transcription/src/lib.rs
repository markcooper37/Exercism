use std::collections::HashMap;

#[derive(Debug, PartialEq)]
pub struct Dna {
    dna: String,
}

#[derive(Debug, PartialEq)]
pub struct Rna {
    rna: String,
}

impl Dna {
    pub fn new(dna: &str) -> Result<Dna, usize> {
        match dna
            .chars()
            .position(|c| c != 'G' && c != 'C' && c != 'T' && c != 'A')
        {
            Some(pos) => Err(pos),
            None => Ok(Dna {
                dna: dna.to_string(),
            }),
        }
    }

    pub fn into_rna(self) -> Rna {
        let dna_nucleotides = vec!['G', 'C', 'T', 'A'];
        let rna_nucleotides = vec!['C', 'G', 'A', 'U'];
        let complements: HashMap<_, _> = dna_nucleotides
            .into_iter()
            .zip(rna_nucleotides.into_iter())
            .collect();
        Rna {
            rna: self
                .dna
                .chars()
                .map(|c| complements.get(&c).unwrap())
                .collect::<String>(),
        }
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        match rna
            .chars()
            .position(|c| c != 'G' && c != 'C' && c != 'U' && c != 'A')
        {
            Some(pos) => Err(pos),
            None => Ok(Rna {
                rna: rna.to_string(),
            }),
        }
    }
}
