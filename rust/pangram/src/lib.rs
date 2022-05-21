/// Determine whether a sentence is a pangram.
pub fn is_pangram(sentence: &str) -> bool {
    let sentence = sentence.to_lowercase();
    ('a'..='z').filter(|c| sentence.contains(*c)).count() == 26
}
