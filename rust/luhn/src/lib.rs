/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let code: String = code.chars().filter(|c| !c.is_whitespace()).collect();
    if code.len() <= 1 {
        return false;
    }
    for char in code.chars() {
        if !char.is_digit(10) {
            return false;
        }
    }
    let sum: u32 = code
        .chars()
        .map(|c| c.to_digit(10).unwrap())
        .rev()
        .enumerate()
        .map(|(i, v)| {
            if i % 2 != 0 {
                ((v * 2) / 10) + ((v * 2) % 10)
            } else {
                v
            }
        })
        .sum();
    sum % 10 == 0
}
