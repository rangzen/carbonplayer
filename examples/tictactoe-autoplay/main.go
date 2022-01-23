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
	p := tictactoe.NewPlayerRandomScore()
	d := decision.NewMinimax(logger, maxPlies)
	node := g.Initial()
	for {
		graphics := [2]string{"O", "X"}
		playerIndex := node.(*tictactoe.Node).PlayerTurn
		fmt.Printf("Player %v (%s) will play on:\n", playerIndex, graphics[playerIndex])
		node.(*tictactoe.Node).ASCIIArt(os.Stdout)
		node = d.NextMove(g, p, node)
		logger.Info("chosen", "base64", node.Base64(), "score", node.Score())
		if g.IsFinal(node) {
			break
		}
	}
	winner := g.WinnerIndex(node)
	logger.Info("finalPosition", "base64", node.Base64(), "winner", winner)
	node.(*tictactoe.Node).ASCIIArt(os.Stdout)
}
