/// Determines whether the supplied string is a valid ISBN number
pub fn is_valid_isbn(isbn: &str) -> bool {
    let mut isbn = isbn.chars().filter(|c| c.is_alphanumeric()).collect::<String>();
    if isbn.len() != 10 {
        return false;
    }
    isbn = isbn.chars().enumerate().filter(|c| c.1.is_digit(10) || c.0 == 9 && c.1 == 'X').map(|c| c.1).collect::<String>();
    if isbn.len() != 10 {
        return false;
    }
    isbn.chars().map(|c| c.to_digit(10).unwrap_or(10)).enumerate().map(|c| c.1 * (10 - c.0 as u32)).sum::<u32>() % 11 == 0
}
