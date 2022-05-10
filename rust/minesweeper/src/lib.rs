pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let minefield: Vec<&[u8]> = minefield.iter().map(|x| x.as_bytes()).collect();
    let mut board: Vec<String> = vec![];
    for (row_index, row) in minefield.iter().enumerate() {
        let mut line = String::new();
        for (column_index, element) in row.iter().enumerate() {
            if *element == '*' as u8 {
                line.push('*');
            } else {
                let mut count = 0;
                for i in -1..=1 {
                    for j in -1..=1 {
                        if (0..(minefield.len())).contains(&(((row_index as i32) + j) as usize))
                            && (0..(row.len())).contains(&(((column_index as i32) + i) as usize))
                            && minefield[((row_index as i32) + j) as usize]
                                [((column_index as i32) + i) as usize]
                                == '*' as u8
                        {
                            count += 1;
                        }
                    }
                }
                if count == 0 {
                    line.push(' ');
                } else {
                    line.push(char::from_digit(count, 10).unwrap());
                }
            }
        }
        board.push(line);
    }
    board
}
