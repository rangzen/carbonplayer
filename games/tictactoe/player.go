package tictactoe

import (
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type player struct{}

func NewPlayer() cp.Player {
	return player{}
}

func (p player) Score(g cp.Game, cpn cp.Node, source bool) float64 {
	var score float64

	final := g.IsFinal(cpn)
	winnerIndex := g.WinnerIndex(cpn)
	if final && winnerIndex != 0 {
		if !source {
			// Final and not a  draw, and it's opp to play: I won
			score = 100
		} else {
			// Final and not a  draw, and it's my turn: I lose
			score = -100
		}
	} else if !final {
		children := g.PossibleChildren(cpn)
		for _, child := range children {
			// Next play lead to a final position and not a draw
			if g.IsFinal(child) && g.WinnerIndex(child) != 0 {
				if !source {
					// It's opp to play, bad news
					score -= 20
				} else {
					// My turn, good news
					score += 20
				}
			}
		}
	}
	return score
}
