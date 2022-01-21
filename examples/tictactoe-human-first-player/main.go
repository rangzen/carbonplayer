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

	g := tictactoe.NewGame()
	p := tictactoe.NewPlayerRandomScore()
	d := decision.NewMinimax(logger, g, maxPlies)
	actualNode := g.Initial()
	for {
		graphics := [2]string{"O", "X"}
		playerIndex := actualNode.(*tictactoe.Node).PlayerTurn
		fmt.Printf("Player %v (%s) will play on:\n", playerIndex, graphics[playerIndex])
		actualNode.(*tictactoe.Node).ASCIIArt(os.Stdout)
		if playerIndex == 0 {
			var move int
			_, err := fmt.Scan(&move)
			if err != nil {
				log.Fatalf("getting human player play: %v", err)
			}
			newNode := actualNode.(*tictactoe.Node).LeafWith(move, 1)
			newNode.PlayerTurn = 1
			actualNode = newNode
		} else {
			actualNode = d.NextMove(g, p, actualNode)
			logger.Info("chosen", "base64", actualNode.Base64(), "score", actualNode.Score())
		}
		if g.IsFinal(actualNode) {
			logger.Info("end", "base64", actualNode.Base64())
			actualNode.(*tictactoe.Node).ASCIIArt(os.Stdout)
			break
		}
	}
}
