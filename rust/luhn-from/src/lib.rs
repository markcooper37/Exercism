use std::fmt::Display;

pub struct Luhn(String);

impl Luhn {
    pub fn is_valid(&self) -> bool {
        self.0
            .chars()
            .filter(|c| !c.is_whitespace())
            .rev()
            .try_fold((0, 0), |(sum, count), c| {
                c.to_digit(10)
                    .map(|n| n + (count % 2) * n)
                    .map(|n| if n > 9 { n - 9 } else { n })
                    .map(|n| (n + sum, count + 1))
            })
            .map_or(false, |(sum, count)| sum % 10 == 0 && count > 1)
    }
}

impl<T: Display> From<T> for Luhn {
    fn from(input: T) -> Self {
        Self(input.to_string())
    }
}
