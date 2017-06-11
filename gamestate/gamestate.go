package gamestate

import (
	"math"
)

type GameState uint32

func New(a, b, c uint8) (gs GameState) {
	gs |= GameState(a)<<16 | GameState(b)<<8 | GameState(c)
	return
}

const AMask = math.MaxUint8 << 16

func (gs GameState) A() (a uint8) {
	a |= uint8((gs & AMask) >> 16)
	return
}

const BMask = math.MaxUint8 << 8

func (gs GameState) B() (b uint8) {
	b |= uint8((gs & BMask) >> 8)
	return
}

const CMask = math.MaxUint8

func (gs GameState) C() (c uint8) {
	c |= uint8(gs & CMask)
	return
}

func (gs GameState) AsSlice() []uint8 {
	return []uint8{gs.A(), gs.B(), gs.C()}
}

func (gs GameState) Magnitude() int {
	return int(gs.A()) + int(gs.B()) + int(gs.C())
}

var TwosMasks = map[string]uint32{
	"A": math.MaxUint8<<16 | math.MaxUint8<<8,
	"B": math.MaxUint8<<8 | math.MaxUint8,
	"C": math.MaxUint8<<16 | math.MaxUint8,
}

func (gs GameState) Twos() map[string]GameState {
	m := make(map[string]GameState)
	for _, key := range []string{"A", "B", "C"} {
		m[key] = GameState(uint32(gs) & TwosMasks[key])
	}
	return m
}
