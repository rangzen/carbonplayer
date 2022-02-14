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

package main

import (
	"fmt"
	"os"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/tictactoe"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
)

var (
	maxPlies = 8
)

func main() {
	stdr.SetVerbosity(3)
	logger := stdr.New(nil)

	g := tictactoe.NewGame()
	p := tictactoe.NewPlayer()
	d := decision.NewMinimax(logger, maxPlies)
	node := g.Initial()
	for {
		graphics := [2]string{"X", "0"}
		playerIndex := node.(*tictactoe.Node).PlayerTurn
		fmt.Printf("Player %v (%s) will play on:\n", playerIndex+1, graphics[playerIndex])
		node.(*tictactoe.Node).ASCIIArt(os.Stdout)
		node = d.NextMove(g, p, node)
		logger.Info("chosen", "string", node.String(), "score", node.Score())
		if g.IsFinal(node) {
			break
		}
	}
	winner := g.WinnerIndex(node)
	logger.Info("finalPosition", "string", node.String(), "winner", winner)
	node.(*tictactoe.Node).ASCIIArt(os.Stdout)
}
