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
	return nm.negamax(node, nm.maxPlies, 1, g, p)
}

// https://en.wikipedia.org/wiki/Negamax#Negamax_base_algorithm
func (nm negamax) negamax(n cp.Node, depth int, color int8, g cp.Game, p cp.Player) cp.Node {
	nm.logger.V(4).Info("entering", "depth", depth, "base64", n.Base64(), "color", color)
	if depth == 0 || g.IsFinal(n) {
		score := p.Score(g, n) * float64(color)
		n.SetScore(score)
		nm.logger.V(6).Info("score", "depth", depth, "base64", n.Base64(), "color", color, "score", score)
		return n
	}
	children := g.PossibleChildren(n)
	value := node.MinusInf
	for _, child := range children {
		v := nm.negamax(child, depth-1, color*-1, g, p)
		if v.Score() > value.Score() {
			value = child
			value.SetScore(v.Score())
			nm.logger.V(6).Info("max", "depth", depth, "base64", child.Base64(), "color", color, "score", v.Score())
		}
	}
	return value
}
