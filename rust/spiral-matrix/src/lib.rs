pub fn spiral_matrix(size: u32) -> Vec<Vec<u32>> {
    let (mut spiral, mut row, mut col, mut num, mut row_change, mut col_change): (Vec<Vec<u32>>, i32, i32, u32, i32, i32) =
        (vec![vec![0; size as usize]; size as usize], 0, 0, 1, 0, 1);
    for _ in 0..size * size {
        (spiral[row as usize][col as usize], num) = (num, num + 1);
        if !(0..size).contains(&((row + row_change) as u32))
            || !(0..size).contains(&((col + col_change) as u32))
            || spiral[(row + row_change) as usize][(col + col_change) as usize] != 0
        {
            (row_change, col_change) = (col_change, -row_change);
        }
        (row, col) = (row + row_change, col + col_change);
    }
    spiral
}
