package gamestate

import (
	"fmt"
	"os"

	"github.com/fighterlyt/permutation"
)

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

var WinningStates = make(map[GameState]bool)

func init() {
	// Fill out WinningStates with permutations of winning_positions
	for _, wp := range winning_positions {
		permutator, err := permutation.NewPerm(wp, nil)
		if err != nil {
			fmt.Printf("Error creating permutator: %s\n", err)
			os.Exit(1)
		}
		for perm, err := permutator.Next(); err == nil; perm, err = permutator.Next() {
			state := New(perm.([]uint8)[0], perm.([]uint8)[1], perm.([]uint8)[2])
			WinningStates[state] = true
		}
	}

	// Add permutations of two equal heaps
	var i uint8
	for i = 1; i <= 15; i++ {
		var seed []uint8 = []uint8{i, i, 0}
		permutator, err := permutation.NewPerm(seed, nil)
		if err != nil {
			fmt.Printf("Error creating permutator: %s\n", err)
			os.Exit(1)
		}
		for perm, err := permutator.Next(); err == nil; perm, err = permutator.Next() {
			state := New(perm.([]uint8)[0], perm.([]uint8)[1], perm.([]uint8)[2])
			WinningStates[state] = true
		}
	}
}

func IsWinningState(gs GameState) bool {
	return WinningStates[gs]
}
