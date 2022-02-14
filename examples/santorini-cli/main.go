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
	"strconv"
	"strings"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
)

var (
	maxPlies = 3
)

func main() {
	stdr.SetVerbosity(3)
	logger := stdr.New(nil)

	g := santorini.NewGame()
	p := santorini.NewPlayer()
	d := decision.NewMinimax(logger, maxPlies)
	actualNode := g.Initial()
	actualNode.(*santorini.Node).Worker[0] = [2]int{1, 3}
	actualNode.(*santorini.Node).Worker[1] = [2]int{3, 3}
	actualNode.(*santorini.Node).Worker[2] = [2]int{1, 1}
	actualNode.(*santorini.Node).Worker[3] = [2]int{3, 1}

	for {
		playerIndex := actualNode.(*santorini.Node).TurnOf
		if playerIndex == 0 {
			actualNode.(*santorini.Node).ASCIIArt(os.Stdout)
			workerIndex := askWorker()
			x, y := askPosition("Move to")
			actualNode.(*santorini.Node).Worker[workerIndex][0] = x
			actualNode.(*santorini.Node).Worker[workerIndex][1] = y
			actualNode.(*santorini.Node).ASCIIArt(os.Stdout)
			if g.IsFinal(actualNode) {
				break
			}
			x, y = askPosition("Build at")
			actualNode.(*santorini.Node).Board[x][y]++
			actualNode.(*santorini.Node).TurnOf = santorini.Player2
			actualNode.(*santorini.Node).ASCIIArt(os.Stdout)
		} else {
			fmt.Println("searching the next move...")
			actualNode = d.NextMove(g, p, actualNode)
			logger.Info("chosen", "string", actualNode.String(), "score", actualNode.Score())
		}
		if g.IsFinal(actualNode) {
			break
		}
	}
	actualNode.(*santorini.Node).ASCIIArt(os.Stdout)
	switch g.WinnerIndex(actualNode) {
	case 1:
		fmt.Println("Winner: Player 1")
	case 2:
		fmt.Println("Winner: Player 2")
	}
}

func askWorker() int {
	fmt.Printf("Worker [x/X]: ")
	var workerMark string
	_, err := fmt.Scan(&workerMark)
	if err != nil {
		log.Fatalf("getting human player worker mark: %v", err)
	}
	var workerIndex int
	if strings.Compare(workerMark, "X") == 0 {
		workerIndex++
	}
	return workerIndex
}

func askPosition(message string) (int, int) {
	fmt.Printf("%s [A1-E5]: ", message)
	var position string
	_, err := fmt.Scan(&position)
	if err != nil {
		log.Fatalf("getting position: %v", err)
	}
	var x, y int
	switch strings.ToLower(position[:1]) {
	case "b":
		x = 1
	case "c":
		x = 2
	case "d":
		x = 3
	case "e":
		x = 4
	}
	y, err = strconv.Atoi(position[1:])
	y--
	if err != nil {
		log.Fatalf("getting Y from position: %v", err)
	}
	return x, y
}
