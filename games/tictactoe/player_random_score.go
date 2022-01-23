package tictactoe

import (
	"math/rand"

	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type playerRandomScore struct{}

func NewPlayerRandomScore() cp.Player {
	return playerRandomScore{}
}

func (p playerRandomScore) Score(g cp.Game, cpn cp.Node, source bool) float64 {
	var score float64
	if g.IsFinal(cpn) {
		if g.WinnerIndex(cpn) == 0 {
			// Draw
			score = 0
		} else {
			// Win
			score = 100

		}
	} else {
		score = rand.Float64()
	}
	if !source {
		score = -1 * score
	}
	return score
}
