package tictactoe_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/go-logr/stdr"
	"github.com/rangzen/carbonplayer/games/tictactoe"
	"github.com/rangzen/carbonplayer/pkg/carbonplayer/decision"
	"github.com/stretchr/testify/assert"
)

func TestNode_String(t *testing.T) {
	type fields struct {
		board [9]uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{board: [9]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0}}, "0........."},
		{"one cross North West", fields{board: [9]uint8{1, 0, 0, 0, 0, 0, 0, 0, 0}}, "0X........"},
		{"one circle South East", fields{board: [9]uint8{0, 0, 0, 0, 0, 0, 0, 0, 2}}, "0........O"},
		{"ending game", fields{board: [9]uint8{1, 1, 2, 0, 1, 1, 2, 2, 2}}, "0XXO.XXOOO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tictactoe.Node{
				Board: tt.fields.board,
			}
			if got := n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_AsciiArt(t *testing.T) {
	type fields struct {
		board [9]uint8
	}
	tests := []struct {
		name   string
		fields fields
		wantW  string
	}{
		{"empty", fields{board: [9]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0}}, "   \n   \n   \n"},
		{"one circle North Ouest", fields{board: [9]uint8{2, 0, 0, 0, 0, 0, 0, 0, 0}}, "O  \n   \n   \n"},
		{"one cross South East", fields{board: [9]uint8{0, 0, 0, 0, 0, 0, 0, 0, 1}}, "   \n   \n  X\n"},
		{"ending game", fields{board: [9]uint8{1, 1, 2, 0, 1, 1, 2, 2, 2}}, "XXO\n XX\nOOO\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tictactoe.Node{
				Board: tt.fields.board,
			}
			w := &bytes.Buffer{}
			n.ASCIIArt(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("AsciiArt() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestNode_LeafWith(t *testing.T) {
	type fields struct {
		board [9]uint8
	}
	type args struct {
		index int
		value uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *tictactoe.Node
	}{
		{"empty to center",
			fields{board: [9]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			args{index: 4, value: 1},
			&tictactoe.Node{Board: [9]uint8{0, 0, 0, 0, 1, 0, 0, 0, 0}},
		},
		{"first index",
			fields{board: [9]uint8{0, 1, 2, 1, 2, 1, 2, 1, 1}},
			args{index: 0, value: 2},
			&tictactoe.Node{Board: [9]uint8{2, 1, 2, 1, 2, 1, 2, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tictactoe.Node{
				Board: tt.fields.board,
			}
			if got := n.LeafWith(tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LeafWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGameMandatoryMoveMinimax tests if the flow is correct with a Minimax.
// OX.    OX.
// .X. -> .X.
// ...    .O.
func TestGameMandatoryMoveMinimaxMaxPlies1(t *testing.T) {
	var maxPlies = 1
	stdr.SetVerbosity(2)
	mm := decision.NewMinimax(logger, maxPlies)
	n := &tictactoe.Node{
		PlayerTurn: 1,
		Board:      [9]uint8{2, 1, 0, 0, 1, 0, 0, 0, 0},
	}

	nextMove := mm.NextMove(g, p, n)

	nextNode := nextMove.(*tictactoe.Node)
	assert.Equal(t, uint8(2), nextNode.Board[7])
}

// TestGameMandatoryMoveMinimax tests if the flow is correct with a Minimax.
// OX.    OX.
// .X. -> .X.
// ...    .O.
func TestGameMandatoryMoveMinimaxMaxPlies2(t *testing.T) {
	var maxPlies = 2
	stdr.SetVerbosity(2)
	mm := decision.NewMinimax(logger, maxPlies)
	n := &tictactoe.Node{
		PlayerTurn: 1,
		Board:      [9]uint8{2, 1, 0, 0, 1, 0, 0, 0, 0},
	}

	nextMove := mm.NextMove(g, p, n)

	nextNode := nextMove.(*tictactoe.Node)
	assert.Equal(t, uint8(2), nextNode.Board[7])
}

// TestGameMandatoryMoveMinimax tests if the flow is correct with a Minimax.
// XO.    XO.
// X.. -> X..
// ...    O..
func TestGameMandatoryMoveMinimaxMaxPlies3(t *testing.T) {
	var maxPlies = 3
	stdr.SetVerbosity(2)
	mm := decision.NewMinimax(logger, maxPlies)
	n := &tictactoe.Node{
		PlayerTurn: 1,
		Board:      [9]uint8{1, 2, 0, 1, 0, 0, 0, 0, 0},
	}

	nextMove := mm.NextMove(g, p, n)

	nextNode := nextMove.(*tictactoe.Node)
	assert.Equal(t, uint8(2), nextNode.Board[6])
}

// TestGameMandatoryMoveMinimax tests if the flow is correct with a Minimax.
// XO.    XO.
// X.. -> X..
// ...    O..
// But in fact, Board[2] will be selected cause with maxPlies high enough
// the algorithm find that the position is a lost position. Any branch will lose.
// The first leaf is selected and nothing win against cause everything
// lead to a loose. The algorithm keep the first space and first choice: 2.
// If the way that the possible leaves are generated, it can be different.
func TestGameMandatoryMoveMinimaxMaxPlies4(t *testing.T) {
	var maxPlies = 4
	stdr.SetVerbosity(2)
	mm := decision.NewMinimax(logger, maxPlies)
	n := &tictactoe.Node{
		PlayerTurn: 1,
		Board:      [9]uint8{1, 2, 0, 1, 0, 0, 0, 0, 0},
	}

	nextMove := mm.NextMove(g, p, n)

	nextNode := nextMove.(*tictactoe.Node)
	assert.Equal(t, uint8(2), nextNode.Board[2])
}
