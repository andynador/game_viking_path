package interfaces

import (
	"github.com/andynador/game_viking_path/app/services/gameContext"
)

type HandlerInterface interface {
	Handle(gameContext *gameContext.GameContext)
}
