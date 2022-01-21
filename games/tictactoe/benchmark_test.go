package tictactoe_test

import (
	"testing"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/tictactoe"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
)

func benchmarkTicTacToeMinimaxAutoplay(maxPlies int, b *testing.B) {
	stdr.SetVerbosity(0)
	logger := stdr.New(nil)
	g := tictactoe.NewGame()
	p := tictactoe.NewPlayerRandomScore()
	d := decision.NewMinimax(logger, g, maxPlies)
	for i := 0; i < b.N; i++ {
		node := g.Initial()
		for {
			node = d.NextMove(g, p, node)
			if g.IsFinal(node) {
				break
			}
		}
	}
}

func BenchmarkTicTacToeMinimaxMaxPlies2(b *testing.B) { benchmarkTicTacToeMinimaxAutoplay(2, b) }
func BenchmarkTicTacToeMinimaxMaxPlies4(b *testing.B) { benchmarkTicTacToeMinimaxAutoplay(4, b) }
func BenchmarkTicTacToeMinimaxMaxPlies6(b *testing.B) { benchmarkTicTacToeMinimaxAutoplay(6, b) }
