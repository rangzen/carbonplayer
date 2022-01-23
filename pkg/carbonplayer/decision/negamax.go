package decision

import (
	"github.com/go-logr/logr"
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/node"
)

type negamax struct {
	logger   logr.Logger
	game     cp.Game
	maxPlies int
}

func NewNegamax(logger logr.Logger, game cp.Game, maxPlies int) *negamax {
	return &negamax{
		logger:   logger.WithName("negamax"),
		game:     game,
		maxPlies: maxPlies,
	}
}

func (nm negamax) NextMove(g cp.Game, p cp.Player, node cp.Node) cp.Node {
	return nm.negamax(node, nm.maxPlies, true, g, p)
}

// https://en.wikipedia.org/wiki/Negamax#Negamax_base_algorithm
func (nm negamax) negamax(n cp.Node, depth int, source bool, g cp.Game, p cp.Player) cp.Node {
	nm.logger.V(4).Info("entering", "depth", depth, "base64", n.Base64())
	if depth == 0 || g.IsFinal(n) {
		score := p.Score(g, n, source)
		n.SetScore(score)
		nm.logger.V(6).Info("score", "depth", depth, "base64", n.Base64(), "score", score)
		return n
	}
	bestNode := node.MinusInf
	children := g.PossibleChildren(n)
	for _, child := range children {
		v := nm.negamax(child, depth-1, !source, g, p)
		if v.Score() > bestNode.Score() {
			bestNode = child
			bestNode.SetScore(v.Score())
			nm.logger.V(6).Info("max", "depth", depth, "base64", child.Base64(), "score", v.Score())
		}
	}
	bestNode.SetScore(-1 * bestNode.Score())
	return bestNode
}
