/*
 * Copyright 2022 Cédric L'HOMME
 *
 * This file is part of the Carbon Player Framework.
 *
 * The Carbon Player Framework is free software: you can redistribute it and/or modify it under the terms of
 * the GNU General Public License as published by the Free Software Foundation, either version 3 of the License,
 * or (at your option) any later version.
 *
 * The Carbon Player Framework is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
 * without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with the Carbon Player Framework.
 * If not, see <https://www.gnu.org/licenses/>.
 */

package decision

import (
	"github.com/go-logr/logr"
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/node"
)

type negamax struct {
	logger   logr.Logger
	maxPlies int
}

func NewNegamax(logger logr.Logger, maxPlies int) *negamax {
	return &negamax{
		logger:   logger.WithName("negamax"),
		maxPlies: maxPlies,
	}
}

func (nm negamax) NextMove(g cp.Game, p cp.Player, node cp.Node) cp.Node {
	return nm.negamax(node, nm.maxPlies, true, g, p)
}

// https://en.wikipedia.org/wiki/Negamax#Negamax_base_algorithm
func (nm negamax) negamax(n cp.Node, depth int, source bool, g cp.Game, p cp.Player) cp.Node {
	nm.logger.V(4).Info("entering", "depth", depth, "string", n.String())
	if depth == 0 || g.IsFinal(n) {
		score := p.Score(g, n, source)
		n.SetScore(score)
		nm.logger.V(6).Info("score", "depth", depth, "string", n.String(), "score", score)
		return n
	}
	bestNode := node.MinusInf
	children := g.PossibleChildren(n)
	for _, child := range children {
		v := nm.negamax(child, depth-1, !source, g, p)
		if v.Score() > bestNode.Score() {
			bestNode = child
			bestNode.SetScore(v.Score())
			nm.logger.V(6).Info("max", "depth", depth, "string", child.String(), "score", v.Score())
		}
	}
	bestNode.SetScore(-1 * bestNode.Score())
	return bestNode
}
