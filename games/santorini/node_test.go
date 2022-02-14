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
	"os"
	"testing"

	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/stretchr/testify/assert"
)

func TestNode_String(t *testing.T) {
	node := santorini.NewNode()
	node.TurnOf = santorini.Player1
	node.Worker = [4]santorini.Position{{1, 3}, {3, 3}, {1, 1}, {3, 1}}
	node.Board = [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{3, 0, 0, 0, 2},
		{0, 2, 0, 0, 3},
		{0, 0, 1, 0, 4},
	}

	assert.Equal(t, "P1.24.44.22.42.00000.00001.30002.02003.00104", node.String())
}

func TestNode_ASCIIArt_RunningBoard(t *testing.T) {
	node := santorini.NewNode()
	node.Worker = [4]santorini.Position{{0, 3}, {1, 2}, {2, 1}, {3, 0}}
	node.Board = [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2},
		{0, 0, 0, 0, 3},
		{0, 0, 0, 0, 4},
	}

	node.ASCIIArt(os.Stdout)
}
