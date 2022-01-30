package santorini_test

import (
	"os"
	"testing"

	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/stretchr/testify/assert"
)

func TestNode_String(t *testing.T) {
	node := santorini.NewNode()
	node.TurnOf = santorini.Player1
	node.Worker = [4]santorini.Position{{1, 3}, {3, 3}, {1, 1}, {3, 1}}
	node.Board = [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{3, 0, 0, 0, 2},
		{0, 2, 0, 0, 3},
		{0, 0, 1, 0, 4},
	}

	assert.Equal(t, "P1.24.44.22.42.00000.00001.30002.02003.00104", node.String())
}

func TestNode_ASCIIArt_RunningBoard(t *testing.T) {
	node := santorini.NewNode()
	node.Worker = [4]santorini.Position{{0, 3}, {1, 2}, {2, 1}, {3, 0}}
	node.Board = [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 2},
		{0, 0, 0, 0, 3},
		{0, 0, 0, 0, 4},
	}

	node.ASCIIArt(os.Stdout)
}
