package tictactoe

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Node struct {
	PlayerTurn uint8
	Board      [9]uint8
	score      float64
}

func NewNode() *Node {
	return &Node{
		PlayerTurn: 0,
		Board:      [9]uint8{},
	}
}

func (n *Node) Score() float64 {
	return n.score
}

func (n *Node) SetScore(s float64) {
	n.score = s
}

func (n *Node) Base64() string {
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(n.PlayerTurn)))
	for _, u := range n.Board {
		switch u {
		case 1:
			sb.WriteByte('X')
		case 2:
			sb.WriteByte('O')
		default:
			sb.WriteByte('.')
		}
	}
	return sb.String()
}

func (n Node) ASCIIArt(w io.Writer) {
	b64 := n.Base64()[1:]
	b := bytes.ReplaceAll([]byte(b64), []byte{'.'}, []byte{' '})
	var output = `%c%c%c
%c%c%c
%c%c%c
`
	_, err := fmt.Fprintf(w, output, b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8])
	if err != nil {
		log.Fatalf("printing ASCII art: %v", err)
	}
}

// LeafWith creates a new Node with a new move.
func (n Node) LeafWith(index int, value uint8) *Node {
	v := Node{Board: n.Board}
	v.Board[index] = value
	return &v
}
