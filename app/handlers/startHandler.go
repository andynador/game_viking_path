package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"github.com/andynador/game_viking_path/app/services/gameContext"
)

const COMMAND_START = "/start"

type StartHandler struct {
	botService *services.BotService
}

func NewStartHandler(botService *services.BotService) *StartHandler {
	return &StartHandler{
		botService: botService,
	}
}

func (handler StartHandler) Handle(gameContext *gameContext.GameContext) {
	handler.botService.Send(
		gameContext.GetUpdate().
			SetText("Привет, Викинг! Для начала, ты можешь осмотреть свой остров /island").
			SetUpdateType(models.MESSAGE_SIMPLE))
}
