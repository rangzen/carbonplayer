/*
 * Copyright 2022 Cédric L'HOMME
 *
 * This file is part of the Carbon Player Framework.
 *
 * The Carbon Player Framework is free software: you can redistribute it and/or modify it under the terms of
 * the GNU General Public License as published by the Free Software Foundation, either version 3 of the License,
 * or (at your option) any later version.
 *
 * The Carbon Player Framework is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
 * without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with the Carbon Player Framework.
 * If not, see <https://www.gnu.org/licenses/>.
 */

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
