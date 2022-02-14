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
