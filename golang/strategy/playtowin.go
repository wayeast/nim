package strategy

import (
	"github.com/wayeast/nim/golang/gamestate"
)

var PlayToWin Strategy = func(gs gamestate.GameState) (gamestate.GameState, error) {
	for ws, mg := range gamestate.WinningStates {
		for key, two := range gs.Twos() {
			if uint32(ws)&gamestate.TwosMasks[key] == uint32(two) && mg < gs.Magnitude() {
				return ws, nil
			}
		}
	}
	return gamestate.GameState(0), gamestate.NoWinningStateError
}
