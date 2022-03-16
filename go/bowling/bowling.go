package bowling

import (
	"errors"
)

// Define the Game type here.
type Game struct {
	frame         int
	rolls         [10][]int
	gameCompleted bool
}

func NewGame() *Game {
	return &Game{frame: 1, rolls: [10][]int{}, gameCompleted: false}
}

func (g *Game) Roll(pins int) error {
	if pins < 0 || pins > 10 {
		return errors.New("number of pins must be between 0 and 10")
	}
	if g.frame < 1 || g.frame > 10 {
		return errors.New("frame number must be between 1 and 10")
	} else if g.frame != 10 {
		if len(g.rolls[g.frame-1]) == 0 {
			g.rolls[g.frame-1] = append(g.rolls[g.frame-1], pins)
			if pins == 10 {
				g.frame++
			}
		} else if len(g.rolls[g.frame-1]) == 1 {
			if g.rolls[g.frame-1][0] == 10 {
				return errors.New("already scored a strike")
			} else if g.rolls[g.frame-1][0]+pins > 10 {
				return errors.New("number of pins in a frame cannot exceed 10")
			} else {
				g.rolls[g.frame-1] = append(g.rolls[g.frame-1], pins)
				g.frame++
			}
		} else {
			return errors.New("frame is already complete")
		}
	} else {
		if len(g.rolls[g.frame-1]) == 0 {
			g.rolls[g.frame-1] = append(g.rolls[g.frame-1], pins)
		} else if len(g.rolls[g.frame-1]) == 1 {
			if g.rolls[g.frame-1][0] == 10 {
				g.rolls[g.frame-1] = append(g.rolls[g.frame-1], pins)
			} else if g.rolls[g.frame-1][0]+pins > 10 {
				return errors.New("number of pins in a frame cannot exceed 10")
			} else {
				g.rolls[g.frame-1] = append(g.rolls[g.frame-1], pins)
				if g.rolls[g.frame-1][0]+pins < 10 {
					g.gameCompleted = true
				}
			}
		} else if len(g.rolls[g.frame-1]) == 2 {
			if g.rolls[g.frame-1][0]+g.rolls[g.frame-1][1] < 10 {
				return errors.New("all rolls have been completed in this frame")
			} else if g.rolls[g.frame-1][0] == 10 && g.rolls[g.frame-1][1] < 10 && g.rolls[g.frame-1][1]+pins > 10 {
				return errors.New("number of pins cannot exceed 10")
			} else {
				g.rolls[g.frame-1] = append(g.rolls[g.frame-1], pins)
				g.gameCompleted = true
			}
		} else {
			return errors.New("all rolls have been completed")
		}
	}
	return nil
}

func (g *Game) Score() (int, error) {
	if !g.gameCompleted {
		return 0, errors.New("game has not been completed")
	}
	strikePreviousFrame := false
	strikeTwoFramesAgo := false
	sparePreviousFrame := false
	score := 0
	for index, frame := range g.rolls {
		if index < 9 {
			if frame[0] == 10 {
				score += 10
				if strikeTwoFramesAgo {
					score += 10
				}
				if strikePreviousFrame || sparePreviousFrame {
					score += 10
				}
				if strikePreviousFrame {
					strikeTwoFramesAgo = true
				}
				strikePreviousFrame = true
				sparePreviousFrame = false
			} else {
				score += frame[0] + frame[1]
				if strikeTwoFramesAgo || sparePreviousFrame {
					score += frame[0]
				}
				if strikePreviousFrame {
					score += frame[0] + frame[1]
				}
				if frame[0]+frame[1] == 10 {
					sparePreviousFrame = true
				}
				strikePreviousFrame = false
				strikeTwoFramesAgo = false
			}
		} else {
			score += frame[0] + frame[1]
			if len(frame) == 3 {
				score += frame[2]
			}
			if strikeTwoFramesAgo || sparePreviousFrame {
				score += frame[0]
			}
			if strikePreviousFrame {
				score += frame[0] + frame[1]
			} 
		}
	}
	return score, nil
}
