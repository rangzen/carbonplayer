package tictactoe

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNode_Base64(t *testing.T) {
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
			n := Node{
				Board: tt.fields.board,
			}
			if got := n.Base64(); got != tt.want {
				t.Errorf("Base64() = %v, want %v", got, tt.want)
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
			n := Node{
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
		want   *Node
	}{
		{"empty to center",
			fields{board: [9]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0}},
			args{index: 4, value: 1},
			&Node{Board: [9]uint8{0, 0, 0, 0, 1, 0, 0, 0, 0}},
		},
		{"first index",
			fields{board: [9]uint8{0, 1, 2, 1, 2, 1, 2, 1, 1}},
			args{index: 0, value: 2},
			&Node{Board: [9]uint8{2, 1, 2, 1, 2, 1, 2, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				Board: tt.fields.board,
			}
			if got := n.LeafWith(tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LeafWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
