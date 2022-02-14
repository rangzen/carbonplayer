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

func (n *Node) String() string {
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
	b64 := n.String()[1:]
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
