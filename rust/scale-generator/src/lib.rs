#[derive(Debug)]
pub enum Error {
    InvalidTonic,
    InvalidInterval,
}

pub struct Scale(Vec<String>);

impl Scale {
    pub fn new(tonic: &str, intervals: &str) -> Result<Scale, Error> {
        let (sharp_scale, flat_scale) = (
            [
                "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#",
            ],
            [
                "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab",
            ],
        );
        let scale = match tonic {
            "C" | "G" | "D" | "A" | "E" | "B" | "F#" | "a" | "e" | "b" | "f#" | "c#" | "g#"
            | "d#" => sharp_scale,
            "F" | "Bb" | "Eb" | "Ab" | "Db" | "Gb" | "d" | "g" | "c" | "f" | "bb" | "eb" => {
                flat_scale
            }
            _ => return Err(Error::InvalidTonic),
        };
        let mut pos = scale
            .iter()
            .position(|&c| c.to_lowercase() == tonic.to_lowercase())
            .unwrap();
        let mut new_scale = vec![scale[pos].to_string()];
        for interval in intervals.chars() {
            pos += match interval {
                'm' => 1,
                'M' => 2,
                'A' => 3,
                _ => return Err(Error::InvalidInterval),
            };
            new_scale.push(scale[pos % 12].to_string())
        }
        Ok(Self(new_scale))
    }

    pub fn chromatic(tonic: &str) -> Result<Scale, Error> {
        Self::new(tonic, "mmmmmmmmmmmm")
    }

    pub fn enumerate(&self) -> Vec<String> {
        self.0.clone()
    }
}
