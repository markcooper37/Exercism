package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int
type Pair [2]int

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
			value, err := strconv.Atoi(element)
			if err != nil {
				return nil, errors.New("invalid input")
			}
			intRow = append(intRow, value)
		}
		matrix = append(matrix, intRow)
	}
	return &matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	var saddlePoints []Pair
	for rowIndex, row := range *m {
		for columnIndex, element := range row {
			biggestInRow := true
			for _, otherElement := range row {
				if otherElement > element {
					biggestInRow = false
				}
			}
			if !biggestInRow {
				continue
			}
			smallestInColumn := true
			for _, newRow := range *m {
				if newRow[columnIndex] < element {
					smallestInColumn = false
				}
			}
			if !smallestInColumn {
				continue
			}
			saddlePoints = append(saddlePoints, Pair{rowIndex, columnIndex})
		}
	}
	return saddlePoints
}
