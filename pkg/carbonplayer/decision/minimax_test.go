package decision_test

import (
	"math"
	"testing"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
	"github.com/stretchr/testify/assert"
)

// https://en.wikipedia.org/wiki/File:Minimax.svgr
func TestWikipediaEnglishExample(t *testing.T) {
	logger := stdr.New(nil)
	g := decision.TreeGame{}
	p := decision.TreePlayer{}
	mm := decision.NewMinimax(logger, g, 4)
	E1 := decision.NewTreeNode("E1")
	E1.SetScore(10)
	E2 := decision.NewTreeNode("E2")
	E2.SetScore(math.Inf(+1))
	D1 := decision.NewTreeNode("D1", E1, E2)
	E3 := decision.NewTreeNode("E3")
	E3.SetScore(5)
	D2 := decision.NewTreeNode("D2", E3)
	C1 := decision.NewTreeNode("C1", D1, D2)
	E4 := decision.NewTreeNode("E4")
	E4.SetScore(-10)
	D3 := decision.NewTreeNode("D3", E4)
	C2 := decision.NewTreeNode("C2", D3)
	B1 := decision.NewTreeNode("B1", C1, C2)
	E5 := decision.NewTreeNode("E5")
	E5.SetScore(7)
	E6 := decision.NewTreeNode("E6")
	E6.SetScore(5)
	D4 := decision.NewTreeNode("D4", E5, E6)
	E7 := decision.NewTreeNode("E7")
	E7.SetScore(math.Inf(-1))
	D5 := decision.NewTreeNode("D5", E7)
	C3 := decision.NewTreeNode("C3", D4, D5)
	E8 := decision.NewTreeNode("E8")
	E8.SetScore(-7)
	E9 := decision.NewTreeNode("E9")
	E9.SetScore(-5)
	D6 := decision.NewTreeNode("D6", E8, E9)
	C4 := decision.NewTreeNode("C4", D6)
	B2 := decision.NewTreeNode("B2", C3, C4)
	A := decision.NewTreeNode("A", B1, B2)

	nm := mm.NextMove(g, p, A)

	assert.Equal(t, "B2:-7.000", nm.Base64())
}

// https://fr.wikipedia.org/wiki/Algorithme_minimax#Exemple
func TestWikipediaFrenchExample(t *testing.T) {
	logger := stdr.New(nil)
	g := decision.TreeGame{}
	p := decision.TreePlayer{}
	mm := decision.NewMinimax(logger, g, 2)
	C1 := decision.NewTreeNode("C1")
	C1.SetScore(12)
	C2 := decision.NewTreeNode("C2")
	C2.SetScore(10)
	C3 := decision.NewTreeNode("C3")
	C3.SetScore(3)
	C4 := decision.NewTreeNode("C4")
	C4.SetScore(5)
	C5 := decision.NewTreeNode("C5")
	C5.SetScore(8)
	C6 := decision.NewTreeNode("C6")
	C6.SetScore(10)
	C7 := decision.NewTreeNode("C7")
	C7.SetScore(11)
	C8 := decision.NewTreeNode("C8")
	C8.SetScore(2)
	C9 := decision.NewTreeNode("C9")
	C9.SetScore(12)
	B1 := decision.NewTreeNode("B1", C1, C2, C3)
	B2 := decision.NewTreeNode("B2", C4, C5, C6)
	B3 := decision.NewTreeNode("B3", C7, C8, C9)
	A := decision.NewTreeNode("A", B1, B2, B3)

	nm := mm.NextMove(g, p, A)

	assert.Equal(t, "B2:5.000", nm.Base64())
}

// https://fr.wikipedia.org/wiki/Algorithme_minimax#Exemple
func TestWikipediaFrenchExamplePartialWithMaxPlies1(t *testing.T) {
	logger := stdr.New(nil)
	g := decision.TreeGame{}
	p := decision.TreePlayer{}
	mm := decision.NewMinimax(logger, g, 1)
	B1 := decision.NewTreeNode("B1")
	B1.SetScore(3)
	B2 := decision.NewTreeNode("B2")
	B2.SetScore(5)
	B3 := decision.NewTreeNode("B3")
	B3.SetScore(2)
	A := decision.NewTreeNode("A", B1, B2, B3)

	nm := mm.NextMove(g, p, A)

	assert.Equal(t, "B2:5.000", nm.Base64())
}
