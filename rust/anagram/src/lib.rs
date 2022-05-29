use std::{collections::HashSet, ops::RangeBounds};

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let word: Vec<char> = word.to_lowercase().chars().collect();
    let possible_anagrams = possible_anagrams
        .iter()
        .map(|w| w.to_lowercase().chars().collect::<Vec<char>>()).filter(|w| *w != word);
    word.sort_unstable();
    possible_anagrams.filter(|w| w.contains(word) && word.contains(w));
    
}
