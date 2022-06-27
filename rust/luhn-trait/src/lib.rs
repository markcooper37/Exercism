use std::fmt::Display;

pub trait Luhn {
    fn valid_luhn(&self) -> bool;
}

impl<T: Display> Luhn for T {
    fn valid_luhn(&self) -> bool {
        self.to_string()
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
