package play

import (
	"fmt"

	"github.com/wayeast/nim/gamestate"
	"github.com/wayeast/nim/strategy"
)

type Player int

const (
	Home    Player = 0
	Visitor Player = 1
)

func ChangePlayer(p Player) Player {
	return (p + 1) % 2
}

func GetNextPlay(p Player, gs gamestate.GameState) gamestate.GameState {
	if p == Home {
		return getHomePlay(gs)
	} else {
		return getVisitorPlay(gs)
	}
}

func getHomePlay(gs gamestate.GameState) gamestate.GameState {
	strategies := []strategy.Strategy{
		strategy.PlayToWin,
		// room for other strategies...
	}

	for _, strategy := range strategies {
		ns, err := strategy(gs)
		if err == nil {
			return ns
		}
	}

	// If all else fails, just play something...
	ns, _ := strategy.JustPlaySomething(gs)
	return ns
}

func getVisitorPlay(gs gamestate.GameState) gamestate.GameState {
	var a, b, c uint8

	fmt.Printf("Current state: %v => Enter A B C: ", gs.AsSlice())
	fmt.Scan(&a, &b, &c)

	for !validInput(gs, a, b, c) {
		fmt.Printf("Current state: %v => Enter A B C: ", gs.AsSlice())
		fmt.Scan(&a, &b, &c)
	}

	return gamestate.New(a, b, c)
}

func validInput(gs gamestate.GameState, a, b, c uint8) bool {
	// if a == 0 && b == 0 && c == 0 {
	// 	//fmt.Println("Not a valid move. Something must be greater than zero.")
	// 	return false
	// }
	if gs == gamestate.New(a, b, c) {
		fmt.Println("Not a valid move. You must take something from one pile.")
		return false
	}
	if a > gs.A() || b > gs.B() || c > gs.C() {
		fmt.Println("Not a valid move. You can't add to any pile.")
		return false
	}
	if a < gs.A() && b < gs.B() {
		fmt.Println("Not a valid move. You may take from only one pile.")
		return false
	}
	if b < gs.B() && c < gs.C() {
		fmt.Println("Not a valid move. You may take from only one pile.")
		return false
	}
	if a < gs.A() && c < gs.C() {
		fmt.Println("Not a valid move. You may take from only one pile.")
		return false
	}
	return true
}
