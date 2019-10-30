package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
)

const COMMAND_ISLAND = "/island"

type IslandHandler struct {
	botService *services.BotService
}

func NewIslandHandler(botService *services.BotService) *IslandHandler {
	return &IslandHandler{
		botService: botService,
	}
}

func (handler IslandHandler) Handle(update *models.Update, user *models.User) {
	handler.botService.Send(
		update.
			SetText("Смотрим остров:").
			SetUpdateType(models.MESSAGE_WITH_KEYBOARD).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton(COMMAND_VIEW_SQUAD),
			)).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton(COMMAND_HIRE_SQUAD),
			)))
}
