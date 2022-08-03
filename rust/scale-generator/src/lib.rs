#[derive(Debug)]
pub enum Error {
    InvalidTonic,
    InvalidIntervalLetter,
    IntervalsTooLong,
    IntervalsTooShort,
}

pub struct Scale(Vec<String>);

impl Scale {
    pub fn new(tonic: &str, intervals: &str) -> Result<Scale, Error> {
        let chromatic: Vec<String>;
        match Scale::chromatic(tonic) {
            Ok(scale) => chromatic = scale.0,
            Err(x) => return Err(x),
        }
        let (mut scale, mut pos) = (vec![chromatic[0].clone()], 0);
        for character in intervals.chars() {
            match character {
                'm' => pos += 1,
                'M' => pos += 2,
                'A' => pos += 3,
                _ => return Err(Error::InvalidIntervalLetter),
            }
            if pos > 12 {
                return Err(Error::IntervalsTooLong);
            }
            scale.push(chromatic[pos].clone());
        }
        match pos {
            12 => Ok(Scale(scale)),
            _ => Err(Error::IntervalsTooShort)
        }
    }

    pub fn chromatic(tonic: &str) -> Result<Scale, Error> {
        let (sharp_scale, flat_scale) = (
            [
                "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#",
            ],
            [
                "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab",
            ],
        );
        let scale: [&str; 12];
        match tonic {
            "C" | "G" | "D" | "A" | "E" | "B" | "F#" | "a" | "e" | "b" | "f#" | "c#" | "g#"
            | "d#" => scale = sharp_scale,
            "F" | "Bb" | "Eb" | "Ab" | "Db" | "Gb" | "d" | "g" | "c" | "f" | "bb" | "eb" => {
                scale = flat_scale
            }
            _ => return Err(Error::InvalidTonic),
        }
        let mut first_position = 0;
        for (index, scale_tonic) in scale.iter().enumerate() {
            if tonic.to_lowercase() == scale_tonic.to_lowercase() {
                first_position = index;
                break;
            }
        }
        let scale = scale.iter().map(|c| c.to_string()).collect::<Vec<String>>();
        let mut chromatic = scale[first_position..].to_vec();
        chromatic.append(&mut scale[..=first_position].to_vec());
        Ok(Scale(chromatic))
    }

    pub fn enumerate(&self) -> Vec<String> {
        self.0.clone()
    }
}
