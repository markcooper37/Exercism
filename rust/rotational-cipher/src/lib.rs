pub fn rotate(input: &str, key: i8) -> String {
    let key = (((key % 26) + 26) % 26) as u8;
    input
        .chars()
        .map(|c| {
            if c.is_ascii_uppercase() {
                (((c as u8 - 65 + key) % 26) + 65) as char
            } else if c.is_ascii_lowercase() {
                (((c as u8 - 97 + key) % 26) + 97) as char
            } else {
                c
            }
        })
        .collect::<String>()
}
