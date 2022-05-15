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
    let mut palindromes: (u64, u64) = (u64::MAX, 0);
    for i in min..=max {
        for j in i..=max {
            if palindromes.0 < i * i {
                break
            }
            match Palindrome::new(i*j) {
                Some(palindrome) => if palindrome.into_inner() < palindromes.0 {
                    palindromes.0 = palindrome.into_inner();
                },
                _ => continue,
            }
        }
    }
    for i in (min..=max).rev() {
        for j in (min..=i).rev() {
            if palindromes.1 > i * i {
                break
            }
            match Palindrome::new(i*j) {
                Some(palindrome) => if palindrome.into_inner() > palindromes.1 {
                    palindromes.1 = palindrome.into_inner();
                },
                _ => continue,
            }
        }
    }
    if palindromes.0 == u64::MAX || (palindromes.1 == 0 && max != 0) {
        None
    } else {
        Some((Palindrome::new(palindromes.0).unwrap(), Palindrome::new(palindromes.1).unwrap()))
    }
}
