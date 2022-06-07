use regex::Regex;
use std::collections::HashMap;

/// Count occurrences of words.
pub fn word_count(words: &str) -> HashMap<String, u32> {
    Regex::new("([a-z]+'[a-z]+)|([a-z]+)|([0-9]+)")
        .unwrap()
        .find_iter(&(words.to_lowercase()))
        .fold(HashMap::<String, u32>::new(), |mut m, x| {
            *m.entry(x.as_str().to_string()).or_default() += 1;
            m
        })
}
