#[derive(Debug, PartialEq)]
pub enum Error {
    SpanTooLong,
    InvalidDigit(char),
}

pub fn lsp(string_digits: &str, span: usize) -> Result<u64, Error> {
    if span > string_digits.len() {
        Err(Error::SpanTooLong)
    } else {
        for character in string_digits.chars() {
            if !character.is_ascii_digit() {
                return Err(Error::InvalidDigit(character));
            }
        } 
        Ok(0)
    } 
}
