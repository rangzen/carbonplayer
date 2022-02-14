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

func TestGoingUp_Value_BothOnTheGround(t *testing.T) {
	var metric santorini.GoingUp
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{1, 3}
	n.Worker[1] = santorini.Position{3, 1}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 3}

	distP1 := metric.Value(n, true)
	distP2 := metric.Value(n, false)

	assert.Equal(t, float64(0), distP1)
	assert.Equal(t, float64(0), distP2)
}

func TestGoingUp_Value_Player1On2(t *testing.T) {
	var metric santorini.GoingUp
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{1, 3}
	n.Worker[1] = santorini.Position{3, 1}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 3}
	n.Board[1][3] = 2

	distP1 := metric.Value(n, true)
	distP2 := metric.Value(n, false)

	assert.Equal(t, float64(.5), distP1)
	assert.Equal(t, float64(-0.5), distP2)
}

// TurnOf Player2
func TestGoingUp_Value_Player1And2On2(t *testing.T) {
	var metric santorini.GoingUp
	n := santorini.NewNode()
	n.TurnOf = santorini.Player2
	n.Worker[0] = santorini.Position{1, 3}
	n.Worker[1] = santorini.Position{3, 1}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 3}
	n.Board[1][3] = 2
	n.Board[1][1] = 2

	distP2 := metric.Value(n, true)
	distP1 := metric.Value(n, false)

	assert.Equal(t, float64(0.5), distP2)
	assert.Equal(t, float64(-0.5), distP1)
}
