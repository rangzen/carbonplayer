package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
)

var (
	maxPlies = 2
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
			fmt.Printf("Worker [1-2]:")
			var workerIndex int
			_, err := fmt.Scan(&workerIndex)
			if err != nil {
				log.Fatalf("getting human player worker index: %v", err)
			}
			fmt.Printf("X [1-5]:")
			var x int
			_, err = fmt.Scan(&x)
			if err != nil {
				log.Fatalf("getting human player worker x: %v", err)
			}
			fmt.Printf("Y [1-5]:")
			var y int
			_, err = fmt.Scan(&y)
			if err != nil {
				log.Fatalf("getting human player worker y: %v", err)
			}
			actualNode.(*santorini.Node).Worker[workerIndex-1][0] = x - 1
			actualNode.(*santorini.Node).Worker[workerIndex-1][1] = y - 1
			actualNode.(*santorini.Node).ASCIIArt(os.Stdout)
			fmt.Println("Build at")
			fmt.Printf("X:")
			_, err = fmt.Scan(&x)
			if err != nil {
				log.Fatalf("getting human player builder x: %v", err)
			}
			fmt.Printf("Y:")
			_, err = fmt.Scan(&y)
			if err != nil {
				log.Fatalf("getting human player builder y: %v", err)
			}
			actualNode.(*santorini.Node).Board[x-1][y-1]++
			actualNode.(*santorini.Node).TurnOf = santorini.Player2
		} else {
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
