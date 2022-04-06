package pascal

func Triangle(n int) [][]int {
	output := [][]int{}
	if n <= 0 {
		return output
	}
	output = append(output, []int{1})
	for i := 1; i < n; i++ {
		row := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, output[i-1][j-1] + output[i-1][j])
			}
		}
		output = append(output, row)
	}
	return output
}
