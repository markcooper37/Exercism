pub fn series(digits: &str, len: usize) -> Vec<String> {
    let mut all_series: Vec<String> = vec![];
    if len > digits.chars().count() {
        return all_series;
    }
    for i in 0..=digits.chars().count()-len {
        all_series.push(digits.chars().skip(i).take(len).collect());
    }
    all_series
}
