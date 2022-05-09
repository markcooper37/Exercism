use std::cmp::Ordering;

#[derive(Debug, PartialEq, Eq)]
pub enum Classification {
    Abundant,
    Perfect,
    Deficient,
}

pub fn classify(num: u64) -> Option<Classification> {
    if num == 0 {
        return None;
    }
    let mut sum = 0;
    for i in 1..=num/2 {
        if num % i == 0 {
            sum += i;
        }
    }
    match sum.cmp(&num) {
        Ordering::Less => Some(Classification::Deficient),
        Ordering::Greater => Some(Classification::Abundant),
        Ordering::Equal => Some(Classification::Perfect)
    }
}
