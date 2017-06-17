package play

import (
	"testing"

	"github.com/wayeast/nim/gamestate"
)

func TestChangePlayer(t *testing.T) {
	var p Player = Home
	if p = ChangePlayer(p); p != Visitor {
		t.Errorf("ChangePlayer failed to change player from Home to Visitor.\n")
	}

	p = Visitor
	if p = ChangePlayer(p); p != Home {
		t.Errorf("ChangePlayer failed to change player from Visitor to Home.\n")
	}
}

func TestValidInput(t *testing.T) {
	var tests = []struct {
		Name     string
		GS       gamestate.GameState
		A, B, C  uint8
		Expected bool
	}{
		{
			Name:     "{gs.A(), gs.B(), gs.C()}",
			GS:       gamestate.New(1, 2, 3),
			A:        1,
			B:        2,
			C:        3,
			Expected: false,
		},
		{
			Name:     "a > gs.A()",
			GS:       gamestate.New(1, 2, 3),
			A:        2,
			B:        2,
			C:        3,
			Expected: false,
		},
		{
			Name:     "a < gs.A() && b < gs.B()",
			GS:       gamestate.New(1, 2, 3),
			A:        0,
			B:        1,
			C:        3,
			Expected: false,
		},
		{
			Name:     "b < gs.B() && c < gs.C()",
			GS:       gamestate.New(1, 2, 3),
			A:        1,
			B:        1,
			C:        2,
			Expected: false,
		},
		{
			Name:     "a < gs.A() && c < gs.C()",
			GS:       gamestate.New(1, 2, 3),
			A:        0,
			B:        2,
			C:        2,
			Expected: false,
		},
		{
			Name:     "Sumpn legit",
			GS:       gamestate.New(1, 2, 3),
			A:        1,
			B:        2,
			C:        1,
			Expected: true,
		},
		{
			Name:     "{0, 0, 1}",
			GS:       gamestate.New(0, 0, 1),
			A:        0,
			B:        0,
			C:        0,
			Expected: true,
		},
	}

	for _, test := range tests {
		check := validInput(test.GS, test.A, test.B, test.C)
		if check != test.Expected {
			t.Errorf("%s: from ABC = %d, %d, %d expected %b, got %b\n",
				test.Name, test.A, test.B, test.C, test.Expected, check)
		}
	}
}
