package santorini_test

import (
	"testing"

	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/stretchr/testify/assert"
)

func TestGoingUp_Value_BothOnTheGround(t *testing.T) {
	var metric santorini.GoingUp
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{1, 3}
	n.Worker[1] = santorini.Position{3, 1}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 3}

	distP1 := metric.Value(n, true)
	distP2 := metric.Value(n, false)

	assert.Equal(t, float64(0), distP1)
	assert.Equal(t, float64(0), distP2)
}

func TestGoingUp_Value_Player1On2(t *testing.T) {
	var metric santorini.GoingUp
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{1, 3}
	n.Worker[1] = santorini.Position{3, 1}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 3}
	n.Board[1][3] = 2

	distP1 := metric.Value(n, true)
	distP2 := metric.Value(n, false)

	assert.Equal(t, float64(.5), distP1)
	assert.Equal(t, float64(-0.5), distP2)
}

// TurnOf Player2
func TestGoingUp_Value_Player1And2On2(t *testing.T) {
	var metric santorini.GoingUp
	n := santorini.NewNode()
	n.TurnOf = santorini.Player2
	n.Worker[0] = santorini.Position{1, 3}
	n.Worker[1] = santorini.Position{3, 1}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 3}
	n.Board[1][3] = 2
	n.Board[1][1] = 2

	distP2 := metric.Value(n, true)
	distP1 := metric.Value(n, false)

	assert.Equal(t, float64(0.5), distP2)
	assert.Equal(t, float64(-0.5), distP1)
}
