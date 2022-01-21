package tictactoe

import cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"

type game struct{}

func NewGame() cp.Game {
	return game{}
}

func (g game) Name() string {
	return "Tic-Tac-Toe"
}

func (g game) Initial() cp.Node {
	return NewNode()
}

func (g game) PlayerIndex(cpn cp.Node) int {
	n := cpn.(*Node)
	return (int(n.PlayerTurn) + 1) % 2

}

func (g game) IsFinal(cpn cp.Node) bool {
	n := cpn.(*Node)
	// Horizontal
	if (n.Board[0] != 0 && n.Board[0] == n.Board[1] && n.Board[1] == n.Board[2]) ||
		(n.Board[3] != 0 && n.Board[3] == n.Board[4] && n.Board[4] == n.Board[5]) ||
		(n.Board[6] != 0 && n.Board[6] == n.Board[7] && n.Board[7] == n.Board[8]) {
		return true
	}
	// Vertical
	if (n.Board[0] != 0 && n.Board[0] == n.Board[3] && n.Board[3] == n.Board[6]) ||
		(n.Board[1] != 0 && n.Board[1] == n.Board[4] && n.Board[4] == n.Board[7]) ||
		(n.Board[2] != 0 && n.Board[2] == n.Board[5] && n.Board[5] == n.Board[8]) {
		return true
	}
	// Diagonal
	if (n.Board[0] != 0 && n.Board[0] == n.Board[4] && n.Board[4] == n.Board[8]) ||
		(n.Board[2] != 0 && n.Board[2] == n.Board[4] && n.Board[4] == n.Board[6]) {
		return true
	}
	// Any empty space?
	for i := 0; i < 9; i++ {
		if n.Board[i] == 0 {
			return false
		}
	}
	return true
}

func (g game) WinnerIndex(cpn cp.Node) int {
	if isWinner(cpn, 1) {
		return 1
	} else if isWinner(cpn, 2) {
		return 2
	} else {
		return 0
	}
}

func isWinner(cpn cp.Node, i uint8) bool {
	n := cpn.(*Node)
	// Horizontal
	if (n.Board[0] == i && n.Board[0] == n.Board[1] && n.Board[1] == n.Board[2]) ||
		(n.Board[3] == i && n.Board[3] == n.Board[4] && n.Board[4] == n.Board[5]) ||
		(n.Board[6] == i && n.Board[6] == n.Board[7] && n.Board[7] == n.Board[8]) {
		return true
	}
	// Vertical
	if (n.Board[0] == i && n.Board[0] == n.Board[3] && n.Board[3] == n.Board[6]) ||
		(n.Board[1] == i && n.Board[1] == n.Board[4] && n.Board[4] == n.Board[7]) ||
		(n.Board[2] == i && n.Board[2] == n.Board[5] && n.Board[5] == n.Board[8]) {
		return true
	}
	// Diagonal
	if (n.Board[0] == i && n.Board[0] == n.Board[4] && n.Board[4] == n.Board[8]) ||
		(n.Board[2] == i && n.Board[2] == n.Board[4] && n.Board[4] == n.Board[6]) {
		return true
	}
	return false
}

func (g game) PossibleChildren(cpn cp.Node) []cp.Node {
	var result []cp.Node
	n := cpn.(*Node)
	for i := 0; i < 9; i++ {
		if n.Board[i] == 0 {
			nextNode := n.LeafWith(i, n.PlayerTurn+1)
			nextNode.PlayerTurn = (n.PlayerTurn + 1) % 2
			result = append(result, nextNode)
		}
	}
	return result
}
