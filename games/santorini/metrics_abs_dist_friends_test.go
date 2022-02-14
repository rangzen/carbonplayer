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

package santorini_test

import (
	"testing"

	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/stretchr/testify/assert"
)

func TestAbsDistToFriends_ValueCenterSameForBothPlayer(t *testing.T) {
	var metric santorini.AbsDistToFriends
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{1, 3}
	n.Worker[1] = santorini.Position{3, 1}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 3}

	distP1 := metric.Value(n, true)
	distP2 := metric.Value(n, false)

	assert.InEpsilon(t, 0.5, distP1, 10e-2)
	assert.InEpsilon(t, 0.5, distP2, 10e-2)
}

func TestAbsDistToFriends_ValueCenterSameForBothPlayerAndNormalizedSoNear1(t *testing.T) {
	var metric santorini.AbsDistToFriends
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{0, 4}
	n.Worker[1] = santorini.Position{4, 0}
	n.Worker[2] = santorini.Position{0, 0}
	n.Worker[3] = santorini.Position{4, 4}

	distP1 := metric.Value(n, true)
	distP2 := metric.Value(n, false)

	assert.InEpsilon(t, 1., distP1, 10e-2)
	assert.InEpsilon(t, 1., distP2, 10e-2)
}
