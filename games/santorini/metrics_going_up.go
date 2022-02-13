package santorini

import (
	"github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type GoingUp struct{}

// Value returns normalized distance between two workers of the same team.
func (m GoingUp) Value(cpn carbonplayer.Node, source bool) float64 {
	n := cpn.(*Node)
	height := n.Board[n.Worker[n.TurnOf*2][0]][n.Worker[n.TurnOf*2][1]]
	height += n.Board[n.Worker[n.TurnOf*2+1][0]][n.Worker[n.TurnOf*2+1][1]]
	// Minimum height for 2 workers: 0, and maximum 4.
	// height of 3 = winning, and is a final state, so no cumulative height of 5 or 6.
	normalizedValue := float64(height) / 4
	if !source {
		normalizedValue *= -1
	}
	return normalizedValue
}
