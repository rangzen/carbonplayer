package santorini

import (
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type player struct{}

func NewPlayer() cp.Player {
	return player{}
}

var metricAbsDistToFriends = AbsDistToFriends{}
var metricAbsDistToEnemies = AbsDistToEnemies{}

func (p player) Score(g cp.Game, cpn cp.Node, itIsYou bool) float64 {
	var score float64
	n := cpn.(*Node)

	final := g.IsFinal(cpn)
	if final {
		if itIsYou {
			return -1000
		} else {
			return 1000
		}
	}

	score += metricAbsDistToFriends.Value(n, itIsYou) * .7
	score += metricAbsDistToEnemies.Value(n, itIsYou) * .3
	if itIsYou {
		score *= -1
	}
	return score
}
