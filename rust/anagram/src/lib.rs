use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let mut anagrams: HashSet<&str> = HashSet::new();
    let word = word.to_lowercase();
    for anagram in possible_anagrams {
        let anagram_lower = anagram.to_lowercase();
        if anagram_lower == word {
            continue;
        }
        let mut word_chars: Vec<char> = word.chars().collect();
        word_chars.sort_unstable();
        let mut anagram_chars: Vec<char> = anagram_lower.chars().collect();
        anagram_chars.sort_unstable();
        if word_chars == anagram_chars {
            anagrams.insert(anagram);
        }
    }
    anagrams
}
