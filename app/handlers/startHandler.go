package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
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

func (handler StartHandler) Handle(update *models.Update, user *models.User) {
	handler.botService.Send(
		update.
			SetText("Привет, Викинг! Для начала, ты можешь осмотреть свой остров /island").
			SetUpdateType(models.MESSAGE_SIMPLE))
}
