use std::cmp::Ordering;

#[derive(Debug, PartialEq, Eq)]
pub enum Classification {
    Abundant,
    Perfect,
    Deficient,
}

pub fn classify(num: u64) -> Option<Classification> {
    if num == 0 {
        None
    } else {
        match (1..=num/2).filter(|&s| num % s == 0).sum::<u64>().cmp(&num) {
            Ordering::Less => Some(Classification::Deficient),
            Ordering::Greater => Some(Classification::Abundant),
            Ordering::Equal => Some(Classification::Perfect),
        }
    }
}
