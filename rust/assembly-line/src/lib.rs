// This stub file contains items which aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub fn production_rate_per_hour(speed: u8) -> f64 {
    if speed <= 4 {
        (speed as u32 * 221) as f64
    } else if speed > 4 && speed <= 8 {
        ((speed as u32 * 221) as f64) * 0.9
    }  else {
        ((speed as u32 * 221) as f64) * 0.77
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) / 60.0) as u32
}
