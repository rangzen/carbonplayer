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

package node

import (
	"math"

	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

var PlusInf cp.Node = plus{}
var MinusInf cp.Node = minus{}

type plus struct {
}

func (p plus) Score() float64 {
	return math.Inf(+1)
}

func (p plus) SetScore(_ float64) {
	panic("should not change this score")
}

func (p plus) String() string {
	return "+inf"
}

type minus struct {
}

func (m minus) Score() float64 {
	return math.Inf(-1)
}

func (m minus) SetScore(_ float64) {
	panic("should not change this score")
}

func (m minus) String() string {
	return "-inf"
}
