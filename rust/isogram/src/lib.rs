use std::collections::HashSet;

pub fn check(candidate: &str) -> bool {
    let mut chars_in_candidate = HashSet::new();
    candidate
        .to_lowercase()
        .chars()
        .filter(|c| c.is_alphabetic())
        .all(|c| chars_in_candidate.insert(c))
}
