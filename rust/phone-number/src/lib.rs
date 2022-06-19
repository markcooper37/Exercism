pub fn number(user_number: &str) -> Option<String> {
    let user_number = user_number.chars().filter(|c| c.is_ascii_digit()).collect::<String>();
    if !(10..=11).contains(&user_number.len()) || user_number.len() == 11 && user_number.chars().nth(0).unwrap() != '1' {
        return None;
    }
    let user_number = user_number.get(user_number.len()-10..).unwrap().to_string();
    if user_number.chars().nth(0).unwrap() <= '1' ||  user_number.chars().nth(3).unwrap() <= '1' {
        return None;
    }
    Some(user_number)
}
