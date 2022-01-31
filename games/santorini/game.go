package santorini

import (
	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

const (
	Player1 = 0
	Player2 = 1
)

type Game struct{}

func NewGame() Game {
	return Game{}
}

func (g Game) Initial() cp.Node {
	return NewNode()
}

func (g Game) PossibleChildren(cpn cp.Node) []cp.Node {
	n := cpn.(*Node)
	switch {
	// First turn, Player 1 choose workers positions
	case n.TurnOf == Player1 &&
		isWorkerPosEqual(n.Worker[0], 0, 0) &&
		isWorkerPosEqual(n.Worker[1], 0, 0):
		result := make([]cp.Node, 0)
		for x1 := 0; x1 < 5; x1++ {
			for y1 := 0; y1 < 5; y1++ {
				for x2 := 0; x2 < 5; x2++ {
					for y2 := 0; y2 < 5; y2++ {
						if x1 == x2 && y1 == y2 {
							continue
						}
						newNode := NewNode()
						newNode.TurnOf = Player2
						newNode.Worker[0] = [2]int{x1, y1}
						newNode.Worker[1] = [2]int{x2, y2}
						result = append(result, newNode)
					}
				}
			}
		}
		return result
	// Second turn, Player 2 choose workers positions
	case n.TurnOf == Player2 &&
		isWorkerPosEqual(n.Worker[2], 0, 0) &&
		isWorkerPosEqual(n.Worker[3], 0, 0):
		result := make([]cp.Node, 0)
		for x1 := 0; x1 < 5; x1++ {
			for y1 := 0; y1 < 5; y1++ {
				if isWorkerPosEqual(n.Worker[0], x1, y1) ||
					isWorkerPosEqual(n.Worker[1], x1, y1) {
					continue
				}
				for x2 := 0; x2 < 5; x2++ {
					for y2 := 0; y2 < 5; y2++ {
						if x1 == x2 && y1 == y2 {
							continue
						}
						if isWorkerPosEqual(n.Worker[0], x2, y2) ||
							isWorkerPosEqual(n.Worker[1], x2, y2) {
							continue
						}
						newNode := NewNode()
						newNode.TurnOf = Player1
						newNode.Worker[2] = [2]int{x1, y1}
						newNode.Worker[3] = [2]int{x2, y2}
						result = append(result, newNode)
					}
				}
			}
		}
		return result
	// Normal turn
	default:
		result := make([]cp.Node, 0)
		for i := 0; i < 2; i++ {
			workerIndex := i
			if n.TurnOf == Player2 {
				workerIndex += 2
			}
			for mx := -1; mx < 2; mx++ {
				for my := -1; my < 2; my++ {
					if mx == 0 && my == 0 {
						continue
					}
					if !n.isWorkerCanMoveRelTo(workerIndex, mx, my) {
						continue
					}
					newPos := [2]int{n.Worker[workerIndex][0] + mx, n.Worker[workerIndex][1] + my}
					nodeAfterMove := n.Clone()
					nodeAfterMove.Worker[workerIndex] = newPos
					// Do not construct if you go to a level 3 building.
					// It is an instant win.
					if nodeAfterMove.Board[newPos[0]][newPos[1]] == 3 {
						result = append(result, nodeAfterMove)
						continue
					}
					for bx := -1; bx < 2; bx++ {
						for by := -1; by < 2; by++ {
							if bx == 0 && by == 0 {
								continue
							}
							if !nodeAfterMove.isWorkerCanBuildRelTo(workerIndex, bx, by) {
								continue
							}
							newNode := nodeAfterMove.Clone()
							newNode.TurnOf = n.NextPlayer()
							newNode.BuildOn(
								nodeAfterMove.Worker[workerIndex][0]+bx,
								nodeAfterMove.Worker[workerIndex][1]+by,
							)
							result = append(result, newNode)
						}
					}
				}
			}
		}
		return result
	}
}

// isWorkerPosEqual returns true if the worker position is equals to x and y.
func isWorkerPosEqual(w [2]int, x, y int) bool {
	return w[0] == x && w[1] == y
}

func (g Game) IsFinal(cpn cp.Node) bool {
	n := cpn.(*Node)
	// A worker is on a 3 level building?
	for i := 0; i < 4; i++ {
		if n.Board[n.Worker[i][0]][n.Worker[i][1]] == 3 {
			return true
		}
	}

	// There is no possible move?
	children := g.PossibleChildren(cpn)
	if len(children) == 0 {
		return true
	}
	return false
}

func (g Game) WinnerIndex(cpn cp.Node) int {
	return (cpn.(*Node).TurnOf + 1) % 2
}
