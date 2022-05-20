pub fn find_saddle_points(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
    let input = input.to_vec();
    let mut saddle_points: Vec<(usize, usize)> = vec![];
    for row in input.iter().enumerate() {
        'column_loop: for column in row.1.iter().enumerate() {
            if column.1 != row.1.iter().max().unwrap() {
                continue
            }
            for i in input.iter() {
                if *column.1 > i[column.0] {
                    continue 'column_loop
                }
            }
            saddle_points.push((row.0, column.0));
        }
    }
    saddle_points
}
