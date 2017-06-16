package strategy

import (
	"github.com/wayeast/nim/gamestate"
)

type Strategy func(gamestate.GameState) (gamestate.GameState, error)
