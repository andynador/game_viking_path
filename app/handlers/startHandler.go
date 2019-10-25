package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
)

type StartHandler struct {
	botService *services.BotService
}

func NewStartHandler(botService *services.BotService) *StartHandler {
	return &StartHandler{
		botService: botService,
	}
}

func (handler StartHandler) Handle(update *models.Update) {
	handler.botService.Send(
		update.
			SetText("Привет, Викинг!").
			SetUpdateType(models.MESSAGE_WITH_KEYBOARD).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton("adsf"),
				models.NewKeyboardButton("a111"),
			)).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton("adsf"),
			)))
}
