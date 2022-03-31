pub fn reply(message: &str) -> &str {
    let contains_letters = message.chars().any(char::is_alphabetic);
    if message.trim().is_empty() {
        "Fine. Be that way!"
    } else if message.trim().chars().last().unwrap() == '?' {
        if message == message.to_ascii_uppercase() && contains_letters {
            "Calm down, I know what I'm doing!"
        } else {
            "Sure."
        }
    } else if message == message.to_ascii_uppercase() && contains_letters {
        "Whoa, chill out!"
    } else {
        "Whatever."
    }
}
