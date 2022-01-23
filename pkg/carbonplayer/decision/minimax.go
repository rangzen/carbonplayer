package decision

import (
	"github.com/go-logr/logr"
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/node"
)

type minimax struct {
	logger   logr.Logger
	game     cp.Game
	maxPlies int
}

func NewMinimax(l logr.Logger, g cp.Game, maxPlies int) *minimax {
	return &minimax{
		logger:   l.WithName("minimax"),
		game:     g,
		maxPlies: maxPlies,
	}
}

func (m minimax) NextMove(g cp.Game, p cp.Player, n cp.Node) cp.Node {
	return m.minimax(n, m.maxPlies, true, g, p)
}

// From https://en.wikipedia.org/wiki/Minimax#Pseudocode with Node
func (m minimax) minimax(n cp.Node, depth int, maximizingPlayer bool, g cp.Game, p cp.Player) cp.Node {
	m.logger.V(4).Info("entering", "depth", depth, "base64", n.Base64(), "maximizingPlayer", maximizingPlayer)
	if depth == 0 || g.IsFinal(n) {
		score := p.Score(g, n, maximizingPlayer)
		n.SetScore(score)
		m.logger.V(6).Info("score", "depth", depth, "base64", n.Base64(), "maximizingPlayer", maximizingPlayer, "score", score)
		return n
	}
	children := g.PossibleChildren(n)
	if maximizingPlayer {
		bestNode := node.MinusInf
		for _, child := range children {
			mm := m.minimax(child, depth-1, false, g, p)
			if mm.Score() > bestNode.Score() {
				bestNode = child
				bestNode.SetScore(mm.Score())
				m.logger.V(6).Info("max", "depth", depth, "base64", child.Base64(), "maximizingPlayer", maximizingPlayer, "score", mm.Score())
			}
		}
		return bestNode
	}
	bestNode := node.PlusInf
	for _, child := range children {
		mm := m.minimax(child, depth-1, true, g, p)
		if mm.Score() < bestNode.Score() {
			bestNode = child
			bestNode.SetScore(mm.Score())
			m.logger.V(6).Info("min", "depth", depth, "base64", child.Base64(), "maximizingPlayer", maximizingPlayer, "score", mm.Score())
		}
	}
	return bestNode
}
