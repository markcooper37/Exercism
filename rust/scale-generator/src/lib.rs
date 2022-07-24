// You should change this.
//
// Depending on your implementation, there are a variety of potential errors
// which might occur. They aren't checked by the test suite in order to
// allow the greatest freedom of implementation, but real libraries should
// provide useful, descriptive errors so that downstream code can react
// appropriately.
//
// One common idiom is to define an Error enum which wraps all potential
// errors. Another common idiom is to use a helper type such as failure::Error
// which does more or less the same thing but automatically.
#[derive(Debug)]
pub enum Error {
    InvalidTonic,
    InvalidIntervals,
}

const SHARP_SCALE: [&str; 12] = [
    "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#",
];
const FLAT_SCALE: [&str; 12] = [
    "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab",
];

pub struct Scale(Vec<String>);

impl Scale {
    pub fn new(tonic: &str, intervals: &str) -> Result<Scale, Error> {
        let chromatic: Vec<String>;
        match Scale::chromatic(tonic) {
            Ok(scale) => chromatic = scale.0,
            Err(x) => return Err(x),
        }
        let mut scale = vec![chromatic[0].clone()];
        let mut pos = 0;
        for character in intervals.chars() {
            match character {
                'm' => {
                    pos += 1;
                }
                'M' => {
                    pos += 2;
                }
                'A' => {
                    pos += 3;
                }
                _ => return Err(Error::InvalidIntervals),
            }
            if pos > 12 {
                return Err(Error::InvalidIntervals);
            }
            scale.push(chromatic[pos].clone());
        }
        if pos < 12 {
            return Err(Error::InvalidIntervals);
        }
        Ok(Scale(scale))
    }

    pub fn chromatic(tonic: &str) -> Result<Scale, Error> {
        let mut scale = vec![];
        match tonic {
            "C" | "G" | "D" | "A" | "E" | "B" | "F#" | "a" | "e" | "b" | "f#" | "c#" | "g#"
            | "d#" => {
                let first_position = find_first_position(
                    &format!(
                        "{}{}",
                        (&tonic[..1].to_string()).to_uppercase(),
                        &tonic[1..]
                    ),
                    true,
                )
                .unwrap();
                for i in 0..13 {
                    scale.push(SHARP_SCALE[(first_position + i) % 12].to_string())
                }
            }
            "F" | "Bb" | "Eb" | "Ab" | "Db" | "Gb" | "d" | "g" | "c" | "f" | "bb" | "eb" => {
                let first_position = find_first_position(
                    &format!(
                        "{}{}",
                        (&tonic[..1].to_string()).to_uppercase(),
                        &tonic[1..]
                    ),
                    false,
                )
                .unwrap();
                for i in 0..13 {
                    scale.push(FLAT_SCALE[(first_position + i) % 12].to_string())
                }
            }
            _ => return Err(Error::InvalidTonic),
        }
        Ok(Scale(scale))
    }

    pub fn enumerate(&self) -> Vec<String> {
        self.0.clone()
    }
}

pub fn find_first_position(input_tonic: &str, sharp: bool) -> Option<usize> {
    let scale: [&'static str; 12];
    if sharp {
        scale = SHARP_SCALE;
    } else {
        scale = FLAT_SCALE;
    }
    for (index, tonic) in scale.iter().enumerate() {
        if input_tonic == *tonic {
            return Some(index);
        }
    }
    None
}
