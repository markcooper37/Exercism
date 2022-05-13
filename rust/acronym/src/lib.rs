pub fn abbreviate(phrase: &str) -> String {
    str::replace(phrase, "-", " ")
        .split_whitespace()
        .map(|c| {
            c.chars()
                .enumerate()
                .zip((String::from(c) + "a").chars().skip(1).enumerate())
                .map(|c| {
                    if c.0.0 == 0 && c.0.1.is_alphabetic() {
                        c.0.1.to_string()
                    } else if !c.0.1.is_uppercase()
                        && c.1.1.is_alphabetic()
                        && c.1.1.is_uppercase()
                    {
                        c.1.1.to_string()
                    } else {
                        String::new()
                    }
                })
                .collect::<String>()
        })
        .collect::<String>()
        .to_uppercase()
}
