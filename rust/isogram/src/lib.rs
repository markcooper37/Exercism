use std::collections::HashMap;

pub fn check(candidate: &str) -> bool {
    let mut chars_in_candidate = HashMap::new();
    candidate
        .to_lowercase()
        .chars()
        .filter(|c| c.is_alphabetic())
        .all(|c| match chars_in_candidate.get(&c) {
            Some(_) => false,
            None => {
                chars_in_candidate.insert(c, true);
                true
            }
        })
}
