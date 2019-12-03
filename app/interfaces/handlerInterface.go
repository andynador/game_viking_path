package interfaces

import (
	"github.com/andynador/game_viking_path/app/models"
)

type HandlerInterface interface {
	Handle(gameContext *models.GameContext)
}
