/// `Palindrome` is a newtype which only exists when the contained value is a palindrome number in base ten.
///
/// A struct with a single field which is used to constrain behavior like this is called a "newtype", and its use is
/// often referred to as the "newtype pattern". This is a fairly common pattern in Rust.
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct Palindrome(u64);

impl Palindrome {
    /// Create a `Palindrome` only if `value` is in fact a palindrome when represented in base ten. Otherwise, `None`.
    pub fn new(value: u64) -> Option<Palindrome> {
        if value.to_string() == value.to_string().chars().rev().collect::<String>() {
            Some(Palindrome(value))
        } else {
            None
        }
    }

    /// Get the value of this palindrome.
    pub fn into_inner(self) -> u64 {
        self.0
    }
}

pub fn palindrome_products(min: u64, max: u64) -> Option<(Palindrome, Palindrome)> {
    let mut palindromes: Vec<Palindrome> = vec![];
    for i in min..=max {
        for j in i..=max {
            match Palindrome::new(i*j) {
                Some(palindrome) => palindromes.push(palindrome),
                _ => continue,
            }
        }
    }
    palindromes.sort_by(|a, b| a.into_inner().cmp(&b.into_inner()));
    if palindromes.len() < 2 {
        None
    } else {
        Some((palindromes[0], palindromes[palindromes.len()-1]))
    }
}
