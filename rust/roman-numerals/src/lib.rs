use std::fmt::{Display, Formatter, Result};

pub struct Roman(String);

impl Display for Roman {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f, "{}", self.0)
    }
}

impl From<u32> for Roman {
    fn from(num: u32) -> Self {
        let digits = vec![
            vec!["", "M", "MM", "MMM"],
            vec!["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"],
            vec!["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"],
            vec!["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"],
        ];
        let num = vec![num / 1000, num % 1000 / 100, num % 100 / 10, num % 10]
            .iter()
            .enumerate()
            .map(|(i, d)| digits[i][*d as usize].to_string())
            .collect::<String>();
        Self(num)
    }
}
