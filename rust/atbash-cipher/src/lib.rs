/// "Encipher" with the Atbash cipher.
pub fn encode(plain: &str) -> String {
    plain
        .to_lowercase()
        .chars()
        .filter(|c| c.is_alphanumeric())
        .map(|c| match c.is_alphabetic() {
            true => (219 - (c as u8)) as char,
            false => c,
        })
        .enumerate()
        .map(|c| match c.0 % 5 {
            4 => {
                let mut result = String::from(c.1);
                result.push(' ');
                result
            }
            _ => String::from(c.1),
        })
        .collect::<String>()
        .trim()
        .to_string()
}

/// "Decipher" with the Atbash cipher.
pub fn decode(cipher: &str) -> String {
    cipher
        .chars()
        .filter(|c| !c.is_whitespace())
        .map(|c| match c.is_alphabetic() {
            true => (219 - (c as u8)) as char,
            false => c,
        })
        .collect::<String>()
}
