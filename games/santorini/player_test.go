package santorini_test

import (
	"testing"

	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/stretchr/testify/assert"
)

func TestPlayer_ScoreP1PoVPlayerP1Win(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player2
	n.Worker[0] = [2]int{0, 4}
	n.Worker[1] = [2]int{3, 3}
	n.Worker[2] = [2]int{1, 1}
	n.Worker[3] = [2]int{3, 1}
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

	assert.Equal(t, float64(1000), score)
}

func TestPlayer_ScoreP2PoVPlayerP2Win(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = [2]int{1, 1}
	n.Worker[1] = [2]int{3, 3}
	n.Worker[2] = [2]int{0, 4}
	n.Worker[3] = [2]int{3, 1}
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

	assert.Equal(t, float64(1000), score)
}

func TestPlayer_ScoreP1PoVPlayerP2Win(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = [2]int{1, 1}
	n.Worker[1] = [2]int{3, 3}
	n.Worker[2] = [2]int{0, 4}
	n.Worker[3] = [2]int{3, 1}
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

	assert.Equal(t, float64(-1000), score)
}

func TestPlayer_ScoreP2PoVPlayerP1Win(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player2
	n.Worker[0] = [2]int{0, 4}
	n.Worker[1] = [2]int{3, 3}
	n.Worker[2] = [2]int{1, 1}
	n.Worker[3] = [2]int{3, 1}
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

	assert.Equal(t, float64(-1000), score)
}
