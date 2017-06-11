package strategy

import (
	"github.com/wayeast/nim/gamestate"
)

func PlayToWin(gs gamestate.GameState) (gamestate.GameState, error) {
	for ws, mg := range gamestate.WinningStates {
		for key, two := range gs.Twos() {
			if uint32(ws)&gamestate.TwosMasks[key] == uint32(two) && mg < gs.Magnitude() {
				return ws, nil
			}
		}
	}
	return gamestate.GameState(0), gamestate.NoWinningStateError
}
