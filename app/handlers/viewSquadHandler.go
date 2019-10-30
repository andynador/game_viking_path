package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
)

const COMMAND_VIEW_SQUAD = "Смотрим дружину"

type ViewSquadHandler struct {
	botService *services.BotService
}

func NewViewSquadHandler(botService *services.BotService) *ViewSquadHandler {
	return &ViewSquadHandler{
		botService: botService,
	}
}

func (handler ViewSquadHandler) Handle(update *models.Update, user *models.User) {
	warriors := user.GetWarriors()
	if len(warriors) == 0 {
		handler.botService.Send(
			update.
				SetText("У тебя пока нет ни одного война").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	var text string
	for _, warrior := range warriors {
		text = text + warrior.GetName() + "\n"
	}

	handler.botService.Send(
		update.
			SetText("Вот твои войны: \n" + text).
			SetUpdateType(models.MESSAGE_SIMPLE))
}
