pub struct PascalsTriangle(u32);

impl PascalsTriangle {
    pub fn new(row_count: u32) -> Self {
        PascalsTriangle(row_count)
    }

    pub fn rows(&self) -> Vec<Vec<u32>> {
        let mut rows: Vec<Vec<u32>> = Vec::new();
        for row_num in 0..self.0 {
            rows.push(
                (0..=row_num)
                    .map(|a| {
                        if a == 0 || a == row_num {
                            1
                        } else {
                            rows[(row_num - 1) as usize][(a - 1) as usize]
                                + rows[(row_num - 1) as usize][a as usize]
                        }
                    })
                    .collect(),
            );
        }
        rows
    }
}
