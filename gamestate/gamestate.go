package gamestate

import (
	"errors"
	"math"
)

var NoWinningStateError error = errors.New("No winning states reachable from here.")

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

var two_element_masks [3]uint32 = [3]uint32{
	math.MaxUint8<<16 | math.MaxUint8<<8,
	math.MaxUint8<<8 | math.MaxUint8,
	math.MaxUint8<<16 | math.MaxUint8,
}

func (gs GameState) Twos() []GameState {
	return []GameState{
		GameState(uint32(gs) & two_element_masks[0]),
		GameState(uint32(gs) & two_element_masks[1]),
		GameState(uint32(gs) & two_element_masks[2]),
	}
}

func (gs GameState) CanGoToWinningState() (GameState, error) {
	for ws, mg := range WinningStates {
		for i, two := range gs.Twos() {
			if uint32(ws)&two_element_masks[i] == uint32(two) && mg < gs.Magnitude() {
				return ws, nil
			}
		}
	}
	return GameState(0), NoWinningStateError
}
