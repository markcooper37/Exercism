pub fn build_proverb(list: &[&str]) -> String {
    if list.len() == 0 {
        return String::from("");
    }
    let mut proverb: String = list.iter()
    .zip(list.iter()
    .skip(1))
    .map(|(a, b)| format!("For want of a {} the {} was lost.\n", a, b))
    .collect::<String>();
    proverb = format!("{}And all for the want of a {}.", proverb, list[0]);
    return proverb
}
