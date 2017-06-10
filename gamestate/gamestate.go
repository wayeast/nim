package gamestate

import "math"

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
