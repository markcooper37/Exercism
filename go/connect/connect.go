package connect

import (
	"errors"
	"strings"
)

func ResultOf(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", errors.New("board has no positions")
	}
	rectangleBoard := [][]string{}
	for _, line := range lines {
		rectangleBoard = append(rectangleBoard, strings.Split(line, ""))
	}
	for index, stone := range rectangleBoard[0] {
		if stone == "O" && RouteExists([][2]int{{0, index}}, rectangleBoard, "O") {
			return "O", nil
		}
	}
	rectangleBoard = FlipBoard(rectangleBoard)
	for index, stone := range rectangleBoard[0] {
		if stone == "X" && RouteExists([][2]int{{0, index}}, rectangleBoard, "X") {
			return "X", nil
		}
	}
	return "", nil
}

func RouteExists(route [][2]int, board [][]string, stoneType string) bool {
	if route[len(route)-1][0] == len(board)-1 {
		return true
	}
	for _, stone := range GetValidNextStones(route, board, stoneType) {
		var routeCopy [][2]int
		routeCopy = append(routeCopy, append(route, stone)...)
		if RouteExists(routeCopy, board, stoneType) {
			return true
		}
	}
	return false
}

func GetValidNextStones(route [][2]int, board [][]string, stoneType string) [][2]int {
	potentialPositions := GetPotentialPositions(route[len(route)-1], board)
	actualPositions := [][2]int{}
	for _, position := range potentialPositions {
		if !ContainedInRoute(route, position) && board[position[0]][position[1]] == stoneType {
			actualPositions = append(actualPositions, position)
		}
	}
	return actualPositions
}

func GetPotentialPositions(lastStone [2]int, board [][]string) [][2]int {
	potentialPositions := [][2]int{}
	if lastStone[1] > 0 {
		potentialPositions = append(potentialPositions, [2]int{lastStone[0], lastStone[1] - 1})
	}
	if lastStone[1] < len(board[lastStone[0]])-1 {
		potentialPositions = append(potentialPositions, [2]int{lastStone[0], lastStone[1] + 1})
	}
	if lastStone[0] > 0 {
		potentialPositions = append(potentialPositions, [2]int{lastStone[0] - 1, lastStone[1]})
	}
	if lastStone[0] < len(board)-1 {
		potentialPositions = append(potentialPositions, [2]int{lastStone[0] + 1, lastStone[1]})
	}
	if lastStone[1] > 0 && lastStone[0] < len(board)-1 {
		potentialPositions = append(potentialPositions, [2]int{lastStone[0] + 1, lastStone[1] - 1})
	}
	if lastStone[0] > 0 && lastStone[1] < len(board[lastStone[0]])-1 {
		potentialPositions = append(potentialPositions, [2]int{lastStone[0] - 1, lastStone[1] + 1})
	}
	return potentialPositions
}

func ContainedInRoute(route [][2]int, position [2]int) bool {
	for _, routePosition := range route {
		if position == routePosition {
			return true
		}
	}
	return false
}

func FlipBoard(board [][]string) [][]string {
	newBoard := [][]string{}
	for columnIndex := range board[0] {
		newRow := []string{}
		for rowIndex := range board {
			newRow = append(newRow, board[rowIndex][columnIndex])
		}
		newBoard = append(newBoard, newRow)
	}
	return newBoard
}
