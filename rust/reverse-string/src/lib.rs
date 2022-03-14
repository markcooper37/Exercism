use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let mut new_string = String::from("");
    for g in (*input).graphemes(true) {
        new_string = format!("{}{}", g, new_string)
    }
    new_string
}
