use std::collections::HashMap;

/// Count occurrences of words.
pub fn word_count(words: &str) -> HashMap<String, u32> {
    words
        .split(|c: char| !c.is_alphanumeric() && c != '\'')
        .filter(|s| !s.is_empty())
        .fold(HashMap::new(), |mut m, x| {
            *m.entry(x.trim_matches('\'').to_lowercase()).or_default() += 1;
            m
        })
}
