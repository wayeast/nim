package gamestate

import (
	"errors"
	"fmt"
	"os"

	"github.com/fighterlyt/permutation"
)

var NoWinningStateError error = errors.New("No winning states reachable from here.")

// From: https://en.wikipedia.org/wiki/Nim#Winning_positions
var winning_positions [][]uint8 = [][]uint8{
	{1, 1, 0},
	{1, 2, 3},
	{1, 4, 5},
	{1, 6, 7},
	{1, 8, 9},
	{2, 4, 6},
	{2, 5, 7},
	{3, 4, 7},
	{3, 5, 6},
	{4, 8, 12},
	{4, 9, 13},
	{5, 8, 13},
	{5, 9, 12},
}

var WinningStates = make(map[GameState]int)

func init() {
	// Fill out WinningStates with permutations of winning_positions
	for _, wp := range winning_positions {
		p, err := permutation.NewPerm(wp, nil)
		if err != nil {
			fmt.Printf("Error creating p: %s\n", err)
			os.Exit(1)
		}
		for res, err := p.Next(); err == nil; res, err = p.Next() {
			gs := New(res.([]uint8)[0], res.([]uint8)[1], res.([]uint8)[2])
			WinningStates[gs] = gs.Magnitude()
		}
	}

	// Add permutations of two equal heaps
	var i uint8
	for i = 1; i <= 15; i++ {
		var seed []uint8 = []uint8{i, i, 0}
		p, err := permutation.NewPerm(seed, nil)
		if err != nil {
			fmt.Printf("Error creating p: %s\n", err)
			os.Exit(1)
		}
		for res, err := p.Next(); err == nil; res, err = p.Next() {
			gs := New(res.([]uint8)[0], res.([]uint8)[1], res.([]uint8)[2])
			WinningStates[gs] = gs.Magnitude()
		}
	}
}

func IsWinningState(gs GameState) bool {
	_, ok := WinningStates[gs]
	return ok
}
