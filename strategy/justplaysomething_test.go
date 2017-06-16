package strategy

import (
	"testing"

	"github.com/wayeast/nim/gamestate"
)

func TestJustPlaySomethingReturnsEventually(t *testing.T) {
	var tests = []struct {
		Name   string
		TestGS gamestate.GameState
	}{
		{
			Name:   "(0, 1, 1)",
			TestGS: gamestate.New(0, 1, 1),
		},
		{
			Name:   "(1, 1, 0)",
			TestGS: gamestate.New(1, 1, 0),
		},
		{
			Name:   "(1, 0, 1)",
			TestGS: gamestate.New(1, 0, 1),
		},
	}

	for _, test := range tests {
		for i := 0; i < 100; i++ {
			next, _ := JustPlaySomething(test.TestGS)
			if next == test.TestGS {
				t.Errorf("Hmm, got same result for input %v\n", test.TestGS.AsSlice())
			}
		}
	}
}
