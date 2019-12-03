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

func (handler IslandHandler) Handle(gameContext *models.GameContext) {
	handler.botService.Send(
		gameContext.GetUpdate().
			SetText("Смотрим остров:").
			SetUpdateType(models.MESSAGE_WITH_KEYBOARD).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton(COMMAND_VIEW_SQUAD),
			)).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton(COMMAND_HIRE_SQUAD),
			)).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton(COMMAND_INVASION),
			)))
}
