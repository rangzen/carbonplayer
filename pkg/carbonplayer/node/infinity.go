package node

import (
	"math"

	cp "github.com/rangzen/carbonplayer/pkg/carbonplayer"
)

var PlusInf cp.Node = plus{}
var MinusInf cp.Node = minus{}

type plus struct {
}

func (p plus) Score() float64 {
	return math.Inf(+1)
}

func (p plus) SetScore(_ float64) {
	panic("should not change this score")
}

func (p plus) String() string {
	return "+inf"
}

type minus struct {
}

func (m minus) Score() float64 {
	return math.Inf(-1)
}

func (m minus) SetScore(_ float64) {
	panic("should not change this score")
}

func (m minus) String() string {
	return "-inf"
}
