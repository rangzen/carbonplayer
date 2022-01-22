package tictactoe_test

import (
	"testing"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/tictactoe"
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
)

var (
	g      = tictactoe.NewGame()
	p      = tictactoe.NewPlayerRandomScore()
	logger = stdr.New(nil)
)

func init() {
	stdr.SetVerbosity(0)
}

func benchmarkTicTacToeAutoplay(d cp.Decision, b *testing.B) {
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

func BenchmarkTicTacToeMinimaxMaxPlies2(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewMinimax(logger, g, 2), b)
}
func BenchmarkTicTacToeMinimaxMaxPlies4(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewMinimax(logger, g, 4), b)
}
func BenchmarkTicTacToeMinimaxMaxPlies6(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewMinimax(logger, g, 6), b)
}
func BenchmarkTicTacToeNegamaxMaxPlies2(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewNegamax(logger, g, 2), b)
}
func BenchmarkTicTacToeNegamaxMaxPlies4(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewNegamax(logger, g, 4), b)
}
func BenchmarkTicTacToeNegamaxMaxPlies6(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewNegamax(logger, g, 6), b)
}
