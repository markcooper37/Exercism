#[derive(Debug, Eq, PartialEq)]
pub enum AffineCipherError {
    NotCoprime(i32),
}

pub fn encode(plaintext: &str, a: i32, b: i32) -> Result<String, AffineCipherError> {
    if a % 2 == 0 || a % 13 == 0 {
        Err(AffineCipherError::NotCoprime(a))
    } else {
        let plaintext = plaintext
            .to_lowercase()
            .chars()
            .filter(|c| c.is_ascii_alphanumeric())
            .map(|c| {
                if c.is_alphabetic() {
                    ((((c as i32 - 97) * a + b).rem_euclid(26) + 97) as u8) as char
                } else {
                    c
                }
            })
            .collect::<Vec<char>>()
            .chunks(5)
            .map(|c| c.iter().collect())
            .collect::<Vec<String>>()
            .join(" ");
        Ok(plaintext)
    }
}

pub fn decode(ciphertext: &str, a: i32, b: i32) -> Result<String, AffineCipherError> {
    if a % 2 == 0 || a % 13 == 0 {
        Err(AffineCipherError::NotCoprime(a))
    } else {
        let mut a_inv = 0;
        for value in 0..26 {
            if value * a % 26 == 1 {
                a_inv = value;
                break;
            }
        }
        let ciphertext = ciphertext
            .chars()
            .filter(|c| !c.is_whitespace())
            .map(|c| {
                if c.is_alphabetic() {
                    ((((c as i32 - 97 - b) * a_inv).rem_euclid(26) + 97) as u8) as char
                } else {
                    c
                }
            })
            .collect();
        Ok(ciphertext)
    }
}
