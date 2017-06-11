package strategy

import (
	"testing"

	"github.com/wayeast/nim/gamestate"
)

func TestPlayToWin(t *testing.T) {
	var tests = []struct {
		Name           string
		GS, ExpectedWM gamestate.GameState
		Error          error
	}{
		{
			Name:       "(1, 2, 0)",
			GS:         gamestate.New(1, 2, 0),
			ExpectedWM: gamestate.New(1, 1, 0),
		},
		{
			Name:       "(1, 2, 4)",
			GS:         gamestate.New(1, 2, 4),
			ExpectedWM: gamestate.New(1, 2, 3),
		},
		{
			Name:       "(5, 2, 3)",
			GS:         gamestate.New(5, 2, 3),
			ExpectedWM: gamestate.New(1, 2, 3),
		},
		{
			Name:       "(1, 4, 3)",
			GS:         gamestate.New(1, 4, 3),
			ExpectedWM: gamestate.New(1, 2, 3),
		},
		{
			Name:       "(2, 4, 8)",
			GS:         gamestate.New(2, 4, 8),
			ExpectedWM: gamestate.New(2, 4, 6),
		},
		{
			Name:       "(3, 4, 6)",
			GS:         gamestate.New(3, 4, 6),
			ExpectedWM: gamestate.New(2, 4, 6),
		},
		{
			Name:       "(5, 8, 12)",
			GS:         gamestate.New(5, 8, 12),
			ExpectedWM: gamestate.New(4, 8, 12),
		},
		{
			Name:       "(0, 1, 0)",
			GS:         gamestate.New(0, 1, 0),
			ExpectedWM: gamestate.New(0, 0, 0),
			Error:      gamestate.NoWinningStateError,
		},
	}

	for _, test := range tests {
		wm, err := PlayToWin(test.GS)
		if err != test.Error {
			t.Errorf("%s: received wrong error '%s' -- expected %s\n", test.Name, err, test.Error)
		}
		if wm != test.ExpectedWM {
			t.Errorf("%s: got unexpected winning state %v -- expected %v\n",
				test.Name, wm.AsSlice(), test.ExpectedWM.AsSlice())
		}
	}

	// Check that calling PlayToWin on any WinningState results in an error
	for ws, _ := range gamestate.WinningStates {
		wm, err := PlayToWin(ws)
		if err != gamestate.NoWinningStateError {
			t.Errorf("%v: received wrong error '%s' -- expected %s\n",
				ws.AsSlice(), err, gamestate.NoWinningStateError)
		}
		if wm != gamestate.New(0, 0, 0) {
			t.Errorf("%v: got unexpected winning state %v -- expected %v\n",
				ws.AsSlice(), wm.AsSlice(), gamestate.New(0, 0, 0).AsSlice())
		}
	}
}
