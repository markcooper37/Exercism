package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (*Matrix, error) {
	stringRows := strings.Split(s, "\n")
	var rowsSeparated [][]string
	for _, row := range stringRows {
		rowSeparated := strings.Split(row, " ")
		rowsSeparated = append(rowsSeparated, rowSeparated)
	}
	var matrix Matrix
	for _, row := range rowsSeparated {
		var intRow []int
		for _, element := range row {
			if element == "" {
				continue
			}
			value, err := strconv.Atoi(element)
			if err != nil {
				return nil, errors.New("invalid input")
			}
			intRow = append(intRow, value)
		}
		matrix = append(matrix, intRow)
		if len(intRow) != len(matrix[0]) {
			return nil, errors.New("lengths of rows do not match")
		}
	}
	return &matrix, nil
}

func (m *Matrix) Cols() [][]int {
	var columns [][]int
	for i := 0; i < len((*m)[0]); i++ {
		var column []int
		for _, row := range *m {
			column = append(column, row[i])
		}
		columns = append(columns, column)
	}
	return columns
}

func (m *Matrix) Rows() [][]int {
	var rows [][]int
	for _, row := range *m {
		var newRow []int
		newRow = append(newRow, row...)
		rows = append(rows, newRow)
	}
	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(*m) || col < 0 || col >= len((*m)[0]) {
		return false
	}
	(*m)[row][col] = val
	return true
}
