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
	"fmt"

	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

// This source file contains the tree structure needed to test decision implementations.

type TreeGame struct{}

func (t TreeGame) Initial() cp.Node {
	return NewTreeNode("A")
}

func (t TreeGame) PossibleChildren(n cp.Node) []cp.Node {
	tn := n.(*TreeNode)
	result := make([]cp.Node, len(tn.Leaves))
	for i, l := range tn.Leaves {
		result[i] = l
	}
	return result
}

func (t TreeGame) IsFinal(n cp.Node) bool {
	tn := n.(*TreeNode)
	return len(tn.Leaves) == 0
}

func (t TreeGame) WinnerIndex(_ cp.Node) int {
	panic("not implemented")
}

type TreeNode struct {
	Name   string
	Leaves []*TreeNode
	score  float64
}

func NewTreeNode(name string, leaves ...*TreeNode) *TreeNode {
	return &TreeNode{
		Name:   name,
		Leaves: leaves,
	}
}

func (t *TreeNode) Score() float64 {
	return t.score
}

func (t *TreeNode) SetScore(s float64) {
	t.score = s
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%s:%.3f", t.Name, t.score)
}

type TreePlayer struct{}

func (t TreePlayer) Score(_ cp.Game, s cp.Node, _ bool) float64 {
	return s.Score()
}
