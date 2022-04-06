package minesweeper

import "strconv"

// Annotate returns an annotated board
func Annotate(board []string) []string {
	annotatedBoard := make([]string, len(board))
	for rowIndex, row := range board {
		for columnIndex, space := range row {
			if space == '*' {
				annotatedBoard[rowIndex] += "*"
				continue
			}
			count := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if rowIndex+i < 0 || rowIndex+i >= len(board) || columnIndex+j < 0 || columnIndex+j >= len(board[0]) {
						continue
					}
					if board[rowIndex+i][columnIndex+j] == '*' {
						count++
					}
				}
			}
			if count > 0 {
				strCount := strconv.Itoa(count)
				annotatedBoard[rowIndex] += strCount
			} else {
				annotatedBoard[rowIndex] += " "
			}
		}
	}
	return annotatedBoard
}
