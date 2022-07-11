package stateoftictactoe

import "errors"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	xCount, oCount := CountPlayer(board, 'X'), CountPlayer(board, 'O')
	if xCount != oCount && xCount != oCount+1 {
		return "", errors.New("invalid combination of X and O")
	}
	xWon, oWon := HasWon(board, 'X'), HasWon(board, 'O')
	if xWon && oWon {
		return "", errors.New("game continued after a player won")
	} else if xWon || oWon {
		return Win, nil
	} else if xCount+oCount == 9 {
		return Draw, nil
	}
	return Ongoing, nil
}

func CountPlayer(board []string, player byte) int {
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == player {
				count++
			}
		}
	}
	return count
}

func HasWon(board []string, player byte) bool {
	for i := 0; i < 3; i++ {
		if (board[i][0] == player && board[i][1] == player && board[i][2] == player) || (board[0][i] == player && board[1][i] == player && board[2][i] == player) {
			return true
		}
	}
	return (board[0][0] == player && board[1][1] == player && board[2][2] == player) || (board[0][2] == player && board[1][1] == player && board[2][0] == player)
}
