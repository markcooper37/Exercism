#[derive(Debug)]
pub struct Duration(f64);

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Self(s as f64)
    }
}

pub trait Planet {
    fn years_during(d: &Duration) -> f64;
}

macro_rules! impl_planet {
    (for $($planet:ident; and $period:expr),+) => {
        $(pub struct $planet;
        impl Planet for $planet {
            fn years_during(d: &Duration) -> f64 {
                d.0 / 31557600. / $period
            }
        })*
    };
}

impl_planet!(for Mercury; and 0.2408467, Venus; and 0.61519726, Earth; and 1., Mars; and 1.8808158, Jupiter; and 11.862615, Saturn; and 29.447498, Uranus; and 84.016846, Neptune; and 164.79132);
