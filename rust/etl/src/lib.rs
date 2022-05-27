use std::collections::BTreeMap;

pub fn transform(h: &BTreeMap<i32, Vec<char>>) -> BTreeMap<char, i32> {
    let mut new_map = BTreeMap::new();
    for value in h {
        for character in value.1 {
            new_map.insert((*character).to_ascii_lowercase(), *value.0);
        }
    }
    new_map
}
