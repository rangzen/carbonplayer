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
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type player struct{}

func NewPlayer() cp.Player {
	return player{}
}

var metricAbsDistToFriends = AbsDistToFriends{}
var metricAbsDistToEnemies = AbsDistToEnemies{}
var metricGoingUp = GoingUp{}

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
	score += metricGoingUp.Value(n, itIsYou) * .6
	if itIsYou {
		score *= -1
	}
	return score
}
