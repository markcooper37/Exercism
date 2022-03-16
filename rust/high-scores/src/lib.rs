#[derive(Debug)]
pub struct HighScores<'a> {
    all_scores: &'a [u32]
}

impl<'a> HighScores<'a> {
    pub fn new(scores: &'a [u32]) -> Self {
        HighScores { all_scores: scores }
    }

    pub fn scores(&self) -> &[u32] {
        self.all_scores
    }

    pub fn latest(&self) -> Option<u32> {
        match self.all_scores.len() {
            0 => None,
            n => Some(self.all_scores[n-1]),
        }
    }

    pub fn personal_best(&self) -> Option<u32> {
        let max_value = self.all_scores.iter().max();
        match max_value {
            Some(n) => Some(*n),
            None => None,
        }
    }

    pub fn personal_top_three(&self) -> Vec<u32> {
        let mut top_three = vec![];
        for score in self.all_scores {
            top_three.push(*score);
            top_three.sort();
            top_three.reverse();
            if top_three.len() > 3 {
                top_three.pop();
            }
        }
        top_three
    }
}
