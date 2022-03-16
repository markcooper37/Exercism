// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut magazine_count = HashMap::new();
    let mut note_count = HashMap::new();
    for word in magazine {
        let count = magazine_count.entry(word).or_insert(0);
        *count += 1;
    }
    for word in note {
        let count = note_count.entry(word).or_insert(0);
        *count += 1;
    }
    for (key, value) in note_count {
        if value > *magazine_count.get(key).unwrap_or(&0) {
            return false;
        }
    }
    true
}
