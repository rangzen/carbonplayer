package santorini_test

import (
	"testing"

	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/stretchr/testify/assert"
)

var metric santorini.AbsDistToFriends

func TestAbsDistToFriends_ValueCenterSameForBothPlayer(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{1, 3}
	n.Worker[1] = santorini.Position{3, 1}
	n.Worker[2] = santorini.Position{1, 1}
	n.Worker[3] = santorini.Position{3, 3}

	distP1 := metric.Value(n, true)
	distP2 := metric.Value(n, false)

	assert.InEpsilon(t, 0.5, distP1, 10e-2)
	assert.InEpsilon(t, -0.5, distP2, 10e-2)
}

func TestAbsDistToFriends_ValueCenterSameForBothPlayerAndNormalizedSoNear1(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = santorini.Position{0, 4}
	n.Worker[1] = santorini.Position{4, 0}
	n.Worker[2] = santorini.Position{0, 0}
	n.Worker[3] = santorini.Position{4, 4}

	distP1 := metric.Value(n, true)
	distP2 := metric.Value(n, false)

	assert.InEpsilon(t, 1., distP1, 10e-2)
	assert.InEpsilon(t, -1., distP2, 10e-2)
}
