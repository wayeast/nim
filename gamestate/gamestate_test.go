package gamestate

import (
	"math"
	"testing"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		Name     string
		Initials [3]uint8
		Test     GameState
	}{
		{
			Name:     "Zeroes",
			Initials: [3]uint8{0, 0, 0},
			Test:     GameState(0),
		},
		{
			Name:     "Ones",
			Initials: [3]uint8{1, 1, 1},
			Test:     GameState(65793),
		},
		{
			Name:     "Maxes",
			Initials: [3]uint8{math.MaxUint8, math.MaxUint8, math.MaxUint8},
			Test:     GameState(math.MaxUint32 >> 8),
		},
	}

	for _, test := range tests {
		gs := New(test.Initials[0], test.Initials[1], test.Initials[2])
		if gs^test.Test != 0 {
			t.Errorf("%s: initialized to %d\n", test.Name, gs|test.Test)
		}
	}
}

func TestABC(t *testing.T) {
	var tests = []struct {
		Name     string
		Initials [3]uint8
	}{
		{
			Name:     "Zeroes",
			Initials: [3]uint8{0, 0, 0},
		},
		{
			Name:     "Maxes",
			Initials: [3]uint8{math.MaxUint8, math.MaxUint8, math.MaxUint8},
		},
		{
			Name:     "123",
			Initials: [3]uint8{1, 2, 3},
		},
	}

	for _, test := range tests {
		gs := New(test.Initials[0], test.Initials[1], test.Initials[2])
		if gs.A() != test.Initials[0] || gs.B() != test.Initials[1] || gs.C() != test.Initials[2] {
			t.Errorf("%s: expecting A: %d, B: %d, C: %d :: got A: %d, B: %d, C: %d\n",
				test.Initials[0], test.Initials[1], test.Initials[2],
				gs.A(), gs.B(), gs.C())
		}
	}
}
