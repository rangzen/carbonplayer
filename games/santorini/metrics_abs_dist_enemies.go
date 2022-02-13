package santorini

import (
	"github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type AbsDistToEnemies struct{}

// Value returns normalized distance between two workers of the same team.
// MetricAbsDistToEnemies returns normalized distance between each couple of adversary workers.
func (m AbsDistToEnemies) Value(cpn carbonplayer.Node, source bool) float64 {
	n := cpn.(*Node)
	dist02 := DistanceBetween(n.Worker[0], n.Worker[2])
	dist03 := DistanceBetween(n.Worker[0], n.Worker[3])
	dist12 := DistanceBetween(n.Worker[1], n.Worker[2])
	dist13 := DistanceBetween(n.Worker[1], n.Worker[3])
	value := NormalizeAbsDist([]float64{dist02, dist03, dist12, dist13})
	if !source {
		value *= -1
	}
	return value
}
