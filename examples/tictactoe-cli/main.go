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
	"log"
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

	graphics := [2]string{"X", "0"}

	g := tictactoe.NewGame()
	p := tictactoe.NewPlayer()
	d := decision.NewMinimax(logger, maxPlies)
	actualNode := g.Initial()
	for {
		playerIndex := actualNode.(*tictactoe.Node).PlayerTurn
		if playerIndex == 0 {
			actualNode.(*tictactoe.Node).ASCIIArt(os.Stdout)
			fmt.Printf("Player %v (%s) play on: ", playerIndex+1, graphics[playerIndex])
			var move int
			_, err := fmt.Scan(&move)
			if err != nil {
				log.Fatalf("getting human player play: %v", err)
			}
			newNode := actualNode.(*tictactoe.Node).LeafWith(move-1, 1)
			newNode.PlayerTurn = 1
			actualNode = newNode
		} else {
			actualNode = d.NextMove(g, p, actualNode)
			logger.Info("chosen", "string", actualNode.String(), "score", actualNode.Score())
		}
		if g.IsFinal(actualNode) {
			break
		}
	}
	actualNode.(*tictactoe.Node).ASCIIArt(os.Stdout)
	switch g.WinnerIndex(actualNode) {
	case 0:
		fmt.Println("Draw...")
	case 1:
		fmt.Println("Winner: Player 1")
	case 2:
		fmt.Println("Winner: Player 2")
	}

}
