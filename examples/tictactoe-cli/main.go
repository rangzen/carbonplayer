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
