package strategy

import (
	"math/rand"

	"github.com/wayeast/nim/golang/gamestate"
)

var JustPlaySomething = func(gs gamestate.GameState) (gamestate.GameState, error) {
	for {
		heap := rand.Int() % 3
		if heap == 0 && gs.A() > 0 {
			return gamestate.New(gs.A()-1, gs.B(), gs.C()), nil
		}

		if heap == 1 && gs.B() > 0 {
			return gamestate.New(gs.A(), gs.B()-1, gs.C()), nil
		}

		if heap == 2 && gs.C() > 0 {
			return gamestate.New(gs.A(), gs.B(), gs.C()-1), nil
		}
	}
	return gs, nil
}
