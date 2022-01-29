package santorini

import (
	"io"
	"strconv"
	"strings"
)

type Position [2]int

type Node struct {
	TurnOf int
	Worker [4]Position
	Board  [5][5]int
	score  float64
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Score() float64 {
	return n.score
}

func (n *Node) SetScore(s float64) {
	n.score = s
}

func (n *Node) String() string {
	sb := strings.Builder{}
	if n.TurnOf == Player1 {
		sb.WriteString("P1.")
	} else {
		sb.WriteString("P2.")
	}
	for i := 0; i < 4; i++ {
		sb.WriteString(strconv.Itoa(n.Worker[i][0] + 1))
		sb.WriteString(strconv.Itoa(n.Worker[i][1] + 1))
		sb.WriteString(".")
	}
	for y := 4; y >= 0; y-- {
		for x := 0; x < 5; x++ {
			sb.WriteString(strconv.Itoa(n.Board[x][y]))
		}
		if y != 0 {
			sb.WriteString(".")
		}
	}
	return sb.String()
}

func (n *Node) ASCIIArt(writer io.Writer) {
	if n.TurnOf == 0 {
		_, _ = writer.Write([]byte("Turn of: player 1 (X)\n"))
	} else {
		_, _ = writer.Write([]byte("Turn of: player 2 (O)\n"))
	}
	for y := 4; y >= 0; y-- {
		_, _ = writer.Write([]byte(strconv.Itoa(y + 1)))
		_, _ = writer.Write([]byte(" "))
		for x := 0; x < 5; x++ {
			// Building
			switch {
			case n.Board[x][y] > 0 && n.Board[x][y] < 4:
				_, _ = writer.Write([]byte(strconv.Itoa(n.Board[x][y])))
			case n.Board[x][y] == 4:
				_, _ = writer.Write([]byte("D"))
			default:
				_, _ = writer.Write([]byte(" "))
			}
			// Player
			switch {
			case n.Worker[0][0] == x && n.Worker[0][1] == y:
				_, _ = writer.Write([]byte("x"))
			case n.Worker[1][0] == x && n.Worker[1][1] == y:
				_, _ = writer.Write([]byte("X"))
			case n.Worker[2][0] == x && n.Worker[2][1] == y:
				_, _ = writer.Write([]byte("o"))
			case n.Worker[3][0] == x && n.Worker[3][1] == y:
				_, _ = writer.Write([]byte("O"))
			default:
				_, _ = writer.Write([]byte(" "))
			}
			// Separator
			if x < 4 {
				_, _ = writer.Write([]byte("|"))
			}
		}
		_, _ = writer.Write([]byte("\n"))
	}
	// Letters
	_, _ = writer.Write([]byte("  A  B  C  D  E\n"))
}

func (n *Node) Clone() *Node {
	return &Node{
		TurnOf: n.TurnOf,
		Worker: n.Worker,
		Board:  n.Board,
	}
}

func (n *Node) NextPlayer() int {
	return (n.TurnOf + 1) % 2
}

func (n *Node) BuildOn(x, y int) {
	n.Board[x][y]++
}

func (n *Node) isWorkerCanMoveRelTo(index int, rx int, ry int) bool {
	newX := n.Worker[index][0] + rx
	if newX < 0 || newX > 4 {
		return false
	}
	newY := n.Worker[index][1] + ry
	if newY < 0 || newY > 4 {
		return false
	}
	if n.Board[newX][newY] > 3 {
		return false
	}
	// Already a worker
	for i := 0; i < 4; i++ {
		if i == index {
			continue
		}
		if n.Worker[i][0] == newX && n.Worker[i][1] == newY {
			return false
		}
	}
	// Too high
	if n.Board[newX][newY]-n.Board[n.Worker[index][0]][n.Worker[index][1]] > 1 {
		return false
	}
	return true
}

func (n *Node) isWorkerCanBuildRelTo(index int, rx int, ry int) bool {
	buildX := n.Worker[index][0] + rx
	if buildX < 0 || buildX > 4 {
		return false
	}
	buildY := n.Worker[index][1] + ry
	if buildY < 0 || buildY > 4 {
		return false
	}
	if n.Board[buildX][buildY] > 3 {
		return false
	}
	for i := 0; i < 4; i++ {
		if i == index {
			continue
		}
		if n.Worker[i][0] == buildX && n.Worker[i][1] == buildY {
			return false
		}
	}
	return true
}
