package santorini

import cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"

type player struct{}

func NewPlayer() cp.Player {
	return player{}
}

func (p player) Score(g cp.Game, cpn cp.Node, _ bool) float64 {
	var score float64

	final := g.IsFinal(cpn)
	if final {
		return 100
	}

	return score
}
