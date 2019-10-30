package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"strconv"
	"strings"
)

const COMMAND_WARRIOR = "/warrior_"

type WarriorHandler struct {
	botService *services.BotService
}

func NewWarriorHandler(botService *services.BotService) *WarriorHandler {
	return &WarriorHandler{
		botService: botService,
	}
}

func (handler WarriorHandler) Handle(update *models.Update, user *models.User) {
	text := strings.Replace(update.GetText(), COMMAND_WARRIOR, "", 1)
	id, _ := strconv.Atoi(text)
	warrior := models.GetWarrior(id)
	if warrior == nil {
		handler.botService.Send(
			update.
				SetText("Этот воин нам неизвестен").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	user.AddWarrior(warrior)

	handler.botService.Send(
		update.
			SetText("Воин " + warrior.GetName() + " нанят!").
			SetUpdateType(models.MESSAGE_SIMPLE))
}
