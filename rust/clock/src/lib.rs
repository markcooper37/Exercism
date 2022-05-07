use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut minutes = minutes;
        let mut hours = hours;
        hours += minutes / 60;
        minutes = minutes % 60;
        if minutes < 0 {
            minutes = minutes + 60;
            hours -= 1;
        };
        hours = ((hours % 24) + 24) % 24;
        Self{hours, minutes}
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Self::new(self.hours, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}
