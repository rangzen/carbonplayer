package santorini_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rangzen/carbonplayer/games/santorini"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer"
	"github.com/stretchr/testify/assert"
)

func printNodes(nodes []carbonplayer.Node) {
	for i, n := range nodes {
		fmt.Println("Node", i)
		n.(*santorini.Node).ASCIIArt(os.Stdout)
	}
}

func TestGame_PossibleChildrenPlayer1ShouldPlaceWorker(t *testing.T) {
	n := santorini.NewNode()
	g := santorini.NewGame()

	nodes := g.PossibleChildren(n)

	assert.Equal(t, 25*24, len(nodes))
}

func TestGame_PossibleChildrenPlayer2ShouldPlaceWorker(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player2
	n.Worker[0] = [2]int{0, 0}
	n.Worker[1] = [2]int{0, 1}
	g := santorini.NewGame()

	nodes := g.PossibleChildren(n)

	assert.Equal(t, 23*22, len(nodes))
}

func TestGame_PossibleChildrenFirstTurn(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = [2]int{1, 3}
	n.Worker[1] = [2]int{3, 3}
	n.Worker[2] = [2]int{1, 1}
	n.Worker[3] = [2]int{3, 1}
	g := santorini.NewGame()

	nodes := g.PossibleChildren(n)

	assert.Equal(t, 80, len(nodes))
}

func TestGame_PossibleChildrenNoMoveWithDome(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = [2]int{1, 3}
	n.Worker[1] = [2]int{3, 3}
	n.Worker[2] = [2]int{1, 1}
	n.Worker[3] = [2]int{3, 1}
	n.Board = [5][5]int{
		{4, 4, 4, 4, 4},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
	}
	g := santorini.NewGame()

	nodes := g.PossibleChildren(n)

	assert.Equal(t, 0, len(nodes))
}

func TestGame_PossibleChildrenNoMoveTooHigh(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = [2]int{1, 3}
	n.Worker[1] = [2]int{3, 3}
	n.Worker[2] = [2]int{1, 1}
	n.Worker[3] = [2]int{3, 1}
	n.Board = [5][5]int{
		{4, 4, 4, 2, 4},
		{4, 2, 4, 0, 2},
		{4, 4, 4, 4, 4},
		{4, 2, 4, 2, 4},
		{4, 4, 4, 4, 4},
	}
	g := santorini.NewGame()

	nodes := g.PossibleChildren(n)
	printNodes(nodes)

	assert.Equal(t, 0, len(nodes))
}

func TestGame_PossibleChildrenOneMoveWithoutConstruction(t *testing.T) {
	n := santorini.NewNode()
	n.TurnOf = santorini.Player1
	n.Worker[0] = [2]int{1, 3}
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

	nodes := g.PossibleChildren(n)
	printNodes(nodes)

	assert.Equal(t, 1, len(nodes))
	// Start position should be 2, not 3 because the win is instant when you move
	// to a level 3 building.
	assert.Equal(t, 2, nodes[0].(*santorini.Node).Board[1][3])
}
