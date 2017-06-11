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
				test.Name,
				test.Initials[0], test.Initials[1], test.Initials[2],
				gs.A(), gs.B(), gs.C())
		}
	}
}

func TestMagnitude(t *testing.T) {
	// A few randomly-chosen magnitude checks
	var tests = []struct {
		Name      string
		Gs        GameState
		Magnitude int
	}{
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
		mg := WinningStates[test.Gs]
		if mg != test.Magnitude {
			t.Errorf("%s: expected WinningStates[] %d; got %d\n", test.Name, test.Magnitude, mg)
		}
		if test.Gs.Magnitude() != test.Magnitude {
			t.Errorf("%s: expected Magnitude() %d; got %d\n", test.Name, test.Magnitude, mg)
		}
	}
}

func TestTwos(t *testing.T) {
	var tests = []struct {
		Name         string
		TestGS       GameState
		ExpectedTwos map[string]GameState
	}{
		{
			Name:   "(1, 1, 1)",
			TestGS: New(1, 1, 1),
			ExpectedTwos: map[string]GameState{
				"A": New(1, 1, 0),
				"B": New(0, 1, 1),
				"C": New(1, 0, 1),
			},
		},
		{
			Name:   "(1, 2, 3)",
			TestGS: New(1, 2, 3),
			ExpectedTwos: map[string]GameState{
				"A": New(1, 2, 0),
				"B": New(0, 2, 3),
				"C": New(1, 0, 3),
			},
		},
		{
			Name:   "(3, 2, 1)",
			TestGS: New(3, 2, 1),
			ExpectedTwos: map[string]GameState{
				"A": New(3, 2, 0),
				"B": New(0, 2, 1),
				"C": New(3, 0, 1),
			},
		},
	}

	for _, test := range tests {
		twos := test.TestGS.Twos()
		for key, two := range twos {
			if two != test.ExpectedTwos[key] {
				t.Errorf("%s: unexpected two %v for %v\n",
					test.Name, two, test.ExpectedTwos[key])
			}
		}
	}
}
