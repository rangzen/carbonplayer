package decision_test

import (
	"testing"

	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
	"github.com/stretchr/testify/assert"
)

func TestTreeGame_Initial(t *testing.T) {
	tree := decision.TreeGame{}

	initial := tree.Initial()
	n := initial.(*decision.TreeNode)

	assert.Equal(t, "A:0.000", initial.Base64())
	assert.Equal(t, 0, len(n.Leaves))
}

func TestTreeGame_PossibleChildren(t *testing.T) {
	tree := decision.TreeGame{}
	B1 := decision.NewTreeNode("B1")
	B2 := decision.NewTreeNode("B2")
	A1 := decision.NewTreeNode("A1", B1, B2)

	children := tree.PossibleChildren(A1)

	assert.Equal(t, cp.Node(B1), children[0])
	assert.Equal(t, cp.Node(B2), children[1])
}

func TestTreeNode_Score(t *testing.T) {
	score := 12.34
	tn := decision.TreeNode{}

	tn.SetScore(score)

	assert.Equal(t, tn.Score(), score)
}

func TestTreePlayer_Score(t *testing.T) {
	score := 12.34
	tn := decision.NewTreeNode("")
	tn.SetScore(score)

	s := decision.TreePlayer{}.Score(decision.TreeGame{}, tn)

	assert.Equal(t, s, score)
}
