package tictactoe

import (
	"math/rand"

	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type playerRandomScore struct{}

func NewPlayerRandomScore() cp.Player {
	return playerRandomScore{}
}

func (p playerRandomScore) Score(g cp.Game, cpn cp.Node) float64 {
	if g.IsFinal(cpn) {
		if g.WinnerIndex(cpn) == 0 {
			// Draw
			return 0
		}
		// Win
		return 100
	}
	return rand.Float64()
}
