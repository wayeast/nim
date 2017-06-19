package strategy

import (
	"github.com/wayeast/nim/golang/gamestate"
)

type Strategy func(gamestate.GameState) (gamestate.GameState, error)
