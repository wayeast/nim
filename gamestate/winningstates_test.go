package gamestate

import (
	//"fmt"
	"testing"
)

func TestWinningStatesInitializedCorrectly(t *testing.T) {
	if len(WinningStates) != 117 {
		t.Errorf("Length mismatch of winning_positions (%d) to WinningStates (%d)\n",
			len(winning_positions), len(WinningStates))
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
