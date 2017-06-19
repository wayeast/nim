package main

import (
	"fmt"

	"github.com/wayeast/nim/golang/gamestate"
	"github.com/wayeast/nim/golang/play"
)

func main() {
	manVsMachine()
}

func manVsMachine() {
	// Setup
	var state = gamestate.StartGame()
	var p play.Player = play.Home

	// Play
	for state != gamestate.GameOver {
		p = play.ChangePlayer(p)
		state = play.GetNextPlay(p, state)
	}

	// Announce outcome
	if p == play.Visitor {
		fmt.Println("Congratulations!")
	} else {
		fmt.Println("You lose!")
	}
}
