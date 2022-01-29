package santorini

import cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"

type player struct{}

func NewPlayer() cp.Player {
	return player{}
}

func (p player) Score(g cp.Game, cpn cp.Node, itIsYou bool) float64 {
	var score float64
	n := cpn.(*Node)

	final := g.IsFinal(cpn)
	if final {
		if n.TurnOf != g.WinnerIndex(cpn) && itIsYou {
			return 1000
		} else {
			return -1000
		}
	}

	return score
}
