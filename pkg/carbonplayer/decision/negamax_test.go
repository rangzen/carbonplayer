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

package decision_test

import (
	"testing"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
	"github.com/stretchr/testify/assert"
)

// http://www.csci.csusb.edu/tongyu/research/icc/09/game-engine.html
func TestNegamax_NextMoveCSCIGameEngineExample(t *testing.T) {
	logger := stdr.New(nil)
	g := decision.TreeGame{}
	p := decision.TreePlayer{}
	mm := decision.NewNegamax(logger, 4)
	C1 := decision.NewTreeNode("C1")
	C1.SetScore(2)
	D1 := decision.NewTreeNode("D1")
	D1.SetScore(5)
	D2 := decision.NewTreeNode("D2")
	D2.SetScore(-3)
	C2 := decision.NewTreeNode("C2", D1, D2)
	B1 := decision.NewTreeNode("B1", C1, C2)
	C3 := decision.NewTreeNode("C3")
	C3.SetScore(-6)
	B2 := decision.NewTreeNode("B2", C3)
	D3 := decision.NewTreeNode("D3")
	D3.SetScore(-2)
	C4 := decision.NewTreeNode("C4", D3)
	C5 := decision.NewTreeNode("C5")
	C5.SetScore(4)
	C6 := decision.NewTreeNode("C6")
	C6.SetScore(8)
	B3 := decision.NewTreeNode("B3", C4, C5, C6)
	A := decision.NewTreeNode("A", B1, B2, B3)

	nm := mm.NextMove(g, p, A)

	assert.Equal(t, "B2:-6.000", nm.String())
}
