package santorini

import (
	"github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type AbsDistToFriends struct{}

// Value returns normalized distance between two workers of the same team.
func (m AbsDistToFriends) Value(cpn carbonplayer.Node, source bool) float64 {
	n := cpn.(*Node)
	playerIndex := PlayerIndex(n.TurnOf, source)
	dist := DistanceBetween(n.Worker[playerIndex*2], n.Worker[playerIndex*2+1])
	value := Normalize([]float64{dist})
	if !source {
		value *= -1
	}
	return value
}
