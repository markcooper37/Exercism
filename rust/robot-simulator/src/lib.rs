// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

#[derive(PartialEq, Eq, Debug)]
pub enum Direction {
    North,
    East,
    South,
    West,
}

pub struct Robot {
    position: (i32, i32),
    direction: Direction,
}

impl Robot {
    pub fn new(x: i32, y: i32, d: Direction) -> Self {
        Self {
            position: (x, y),
            direction: d,
        }
    }

    #[must_use]
    pub fn turn_right(self) -> Self {
        match self.direction {
            Direction::North => Self {
                direction: Direction::East,
                ..self
            },
            Direction::East => Self {
                direction: Direction::South,
                ..self
            },
            Direction::South => Self {
                direction: Direction::West,
                ..self
            },
            Direction::West => Self {
                direction: Direction::North,
                ..self
            },
        }
    }

    #[must_use]
    pub fn turn_left(self) -> Self {
        match self.direction {
            Direction::North => Self {
                direction: Direction::West,
                ..self
            },
            Direction::East => Self {
                direction: Direction::North,
                ..self
            },
            Direction::South => Self {
                direction: Direction::East,
                ..self
            },
            Direction::West => Self {
                direction: Direction::South,
                ..self
            },
        }
    }

    #[must_use]
    pub fn advance(self) -> Self {
        match self.direction {
            Direction::North => Self {
                position: (self.position.0, self.position.1 + 1),
                ..self
            },
            Direction::East => Self {
                position: (self.position.0 + 1, self.position.1),
                ..self
            },
            Direction::South => Self {
                position: (self.position.0, self.position.1 - 1),
                ..self
            },
            Direction::West => Self {
                position: (self.position.0 - 1, self.position.1),
                ..self
            },
        }
    }

    #[must_use]
    pub fn instructions(self, instructions: &str) -> Self {
        let mut robot = self;
        for instruction in instructions.chars() {
            match instruction {
                'R' => robot = robot.turn_right(),
                'L' => robot = robot.turn_left(),
                'A' => robot = robot.advance(),
                _ => (),
            }
        }
        robot
    }

    pub fn position(&self) -> (i32, i32) {
        self.position
    }

    pub fn direction(&self) -> &Direction {
        &self.direction
    }
}
