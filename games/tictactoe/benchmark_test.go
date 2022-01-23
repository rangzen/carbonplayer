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
	p      = tictactoe.NewPlayer()
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
	benchmarkTicTacToeAutoplay(decision.NewMinimax(logger, 2), b)
}
func BenchmarkTicTacToeMinimaxMaxPlies4(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewMinimax(logger, 4), b)
}
func BenchmarkTicTacToeMinimaxMaxPlies6(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewMinimax(logger, 6), b)
}
func BenchmarkTicTacToeNegamaxMaxPlies2(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewNegamax(logger, 2), b)
}
func BenchmarkTicTacToeNegamaxMaxPlies4(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewNegamax(logger, 4), b)
}
func BenchmarkTicTacToeNegamaxMaxPlies6(b *testing.B) {
	benchmarkTicTacToeAutoplay(decision.NewNegamax(logger, 6), b)
}
