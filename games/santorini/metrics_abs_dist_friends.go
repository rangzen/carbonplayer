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

type AbsDistToFriends struct{}

// Value returns normalized distance between two workers of the same team.
func (m AbsDistToFriends) Value(cpn carbonplayer.Node, source bool) float64 {
	n := cpn.(*Node)
	playerIndex := PlayerIndex(n.TurnOf, source)
	dist := DistanceBetween(n.Worker[playerIndex*2], n.Worker[playerIndex*2+1])
	value := NormalizeAbsDist([]float64{dist})
	return value
}
