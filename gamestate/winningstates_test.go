package gamestate

import (
	//"fmt"
	"testing"
)

func TestWinningStatesInitializedCorrectly(t *testing.T) {
	// Right number of states
	if len(WinningStates) != 118 {
		t.Errorf("Length mismatch of winning_positions (%d) to WinningStates (%d)\n",
			len(winning_positions), len(WinningStates))
	}

	// A few randomly-chosen magnitude checks
	var tests = []struct {
		Name      string
		Gs        GameState
		Magnitude int
	}{
		{
			Name:      "0, 0, 0",
			Gs:        New(0, 0, 0),
			Magnitude: 0,
		},
		{
			Name:      "1, 1, 0",
			Gs:        New(1, 1, 0),
			Magnitude: 2,
		},
		{
			Name:      "1, 2, 3",
			Gs:        New(1, 2, 3),
			Magnitude: 6,
		},
		{
			Name:      "3, 4, 7",
			Gs:        New(3, 4, 7),
			Magnitude: 14,
		},
		{
			Name:      "4, 9, 13",
			Gs:        New(4, 9, 13),
			Magnitude: 26,
		},
	}

	for _, test := range tests {
		val, ok := WinningStates[test.Gs]
		if !ok || val != test.Magnitude {
			t.Errorf("%s: expected Magnitude %d; got %d\n", test.Magnitude, val)
		}
	}
}

func TestIsWinningState(t *testing.T) {
	for _, wp := range winning_positions {
		gs := New(wp[0], wp[1], wp[2])
		if !IsWinningState(gs) {
			t.Errorf("Winning position %v not found in WinningStates.\n", wp)
		}
	}
}
