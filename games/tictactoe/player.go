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

package tictactoe

import (
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

type player struct{}

func NewPlayer() cp.Player {
	return player{}
}

func (p player) Score(g cp.Game, cpn cp.Node, source bool) float64 {
	var score float64

	final := g.IsFinal(cpn)
	winnerIndex := g.WinnerIndex(cpn)
	if final && winnerIndex != 0 {
		if !source {
			// Final and not a  draw, and it's opp to play: I won
			score = 100
		} else {
			// Final and not a  draw, and it's my turn: I lose
			score = -100
		}
	} else if !final {
		children := g.PossibleChildren(cpn)
		for _, child := range children {
			// Next play lead to a final position and not a draw
			if g.IsFinal(child) && g.WinnerIndex(child) != 0 {
				if !source {
					// It's opp to play, bad news
					score -= 20
				} else {
					// My turn, good news
					score += 20
				}
			}
		}
	}
	return score
}
