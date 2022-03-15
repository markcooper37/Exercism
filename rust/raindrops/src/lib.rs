pub fn raindrops(n: u32) -> String {
    let mut output = String::from("");
    if n % 3 == 0 {
        output = format!("{}{}", output, "Pling")
    }
    if n % 5 == 0 {
        output = format!("{}{}", output, "Plang")
    }
    if n % 7 == 0 {
        output = format!("{}{}", output, "Plong")
    }
    if output == "" {
        return n.to_string();
    }
    output
}
