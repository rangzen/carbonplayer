package santorini_test

import (
	"testing"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
	"github.com/stretchr/testify/assert"
)

func TestPlayer_ScoreP1PoVPlayerP1Win(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player2
	n.Worker[0] = santorini.Position{0, 4}
	n.Worker[1] = santorini.Position{3, 3}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 1}
	n.Board = [5][5]int{
		{4, 4, 4, 4, 3},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
	}
	g := santorini.NewGame()
	p := santorini.NewPlayer()

	score := p.Score(g, n, true)

	assert.Equal(t, float64(-1000), score)
}

func TestPlayer_ScoreP2PoVPlayerP2Win(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{1, 1}
	n.Worker[1] = santorini.Position{3, 3}
	n.Worker[2] = santorini.Position{0, 4}
	n.Worker[3] = santorini.Position{3, 1}
	n.Board = [5][5]int{
		{4, 4, 4, 4, 3},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
	}
	g := santorini.NewGame()
	p := santorini.NewPlayer()

	score := p.Score(g, n, true)

	assert.Equal(t, float64(-1000), score)
}

func TestPlayer_ScoreP1PoVPlayerP2Win(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{1, 1}
	n.Worker[1] = santorini.Position{3, 3}
	n.Worker[2] = santorini.Position{0, 4}
	n.Worker[3] = santorini.Position{3, 1}
	n.Board = [5][5]int{
		{4, 4, 4, 4, 3},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
	}
	g := santorini.NewGame()
	p := santorini.NewPlayer()

	score := p.Score(g, n, false)

	assert.Equal(t, float64(1000), score)
}

func TestPlayer_ScoreP2PoVPlayerP1Win(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player2
	n.Worker[0] = santorini.Position{0, 4}
	n.Worker[1] = santorini.Position{3, 3}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 1}
	n.Board = [5][5]int{
		{4, 4, 4, 4, 3},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
	}
	g := santorini.NewGame()
	p := santorini.NewPlayer()

	score := p.Score(g, n, false)

	assert.Equal(t, float64(1000), score)
}

func TestPlayer_ScoreP2PoVWinIn2MovesWithAMinimaxAndMaxPlies3(t *testing.T) {
	g := santorini.NewGame()
	p := santorini.NewPlayer()
	n := santorini.NewNode()
	n.TurnOf = santorini.Player2
	// Blocked at the left.
	n.Worker[0] = santorini.Position{0, 2}
	// Blocked in the bottom right corner but can move.
	n.Worker[1] = santorini.Position{3, 0}
	// Worker 2 should go to D3
	// then build at E4
	// then worker 3 will go the new level 3 building and win!
	n.Worker[2] = santorini.Position{2, 2}
	n.Worker[3] = santorini.Position{3, 4}
	n.Board = [5][5]int{
		{0, 4, 0, 4, 0},
		{4, 4, 4, 4, 4},
		{4, 4, 0, 4, 2},
		{0, 4, 0, 4, 2},
		{0, 4, 0, 2, 2},
	}
	stdr.SetVerbosity(3)
	logger := stdr.New(nil)
	d := decision.NewMinimax(logger, 3)

	nextMove := d.NextMove(g, p, n)

	nm := nextMove.(*santorini.Node)
	// Should find      P1.13.41.43.45.04040.44444.44042.04042.04032
	// Then             P2.13.51.43.45.04040.44444.44042.14042.04032 (only one possibility)
	// Winning position P1.13.51.43.54.04040.44444.44042.14042.04032
	assert.Equal(t, 3, nm.Worker[2][0])
	assert.Equal(t, 2, nm.Worker[2][1])
	assert.Equal(t, 3, nm.Board[4][3])
}
