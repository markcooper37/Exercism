package spiralmatrix

func SpiralMatrix(size int) [][]int {
	output := make([][]int, size)
	for i := range output {
		output[i] = make([]int, size)
	}
	row, column := 0, 0
	direction := 0
	for i := 1; i <= size*size; i++ {
		output[row][column] = i
		if direction == 0 {
			if column == size-1 || output[row][column+1] != 0 {
				direction++
				row++
			} else {
				column++
			}
		} else if direction == 1 {
			if row == size-1 || output[row+1][column] != 0 {
				direction++
				column--
			} else {
				row++
			}
		} else if direction == 2 {
			if column == 0 || output[row][column-1] != 0 {
				direction++
				row--
			} else {
				column--
			}
		} else {
			if row == 0 || output[row-1][column] != 0 {
				direction = 0
				column++
			} else {
				row--
			}
		}
	}
	return output
}
