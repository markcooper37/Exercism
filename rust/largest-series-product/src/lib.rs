#[derive(Debug, PartialEq)]
pub enum Error {
    SpanTooLong,
    InvalidDigit(char),
}

pub fn lsp(string_digits: &str, span: usize) -> Result<u64, Error> {
    for character in string_digits.chars() {
        if !character.is_ascii_digit() {
            return Err(Error::InvalidDigit(character));
        }
    }
    if span > string_digits.len() {
        Err(Error::SpanTooLong)
    } else {
        if span == 0 {
            return Ok(1);
        }
        let mut something = string_digits
            .bytes()
            .map(|c| (c - 48) as u64)
            .collect::<Vec<u64>>()
            .windows(span)
            .map(|x| x.iter().product())
            .collect::<Vec<u64>>();
        something.sort_by(|a, b| b.cmp(a));
        Ok(something[0])
    }
}
