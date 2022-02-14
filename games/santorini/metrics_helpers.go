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
	"math"
)

// DistanceBetween calculates the distance between two positions.
func DistanceBetween(p1 Position, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64(p1[0]-p2[0]), 2) + math.Pow(float64(p1[1]-p2[1]), 2))
}

// NormalizeAbsDist normalizes absolute distances on the board.
// Max distance on a Santorini board is the diagonal with Sqrt(4^2+4^2) ~ 5.6569
func NormalizeAbsDist(distances []float64) float64 {
	var result float64
	for _, distance := range distances {
		result += distance
	}
	return result / (float64(len(distances)) * 5.657)
}

// PlayerIndex returns the index of the player point of view for scoring
func PlayerIndex(turnOf int, source bool) int {
	if source {
		return (turnOf + 1) % 2
	}
	return turnOf
}
