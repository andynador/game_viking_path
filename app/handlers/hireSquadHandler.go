package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"strconv"
)

const COMMAND_HIRE_SQUAD = "Идём в таверну за набором воинов"

type HireSquadHandler struct {
	botService *services.BotService
}

func NewHireSquadHandler(botService *services.BotService) *HireSquadHandler {
	return &HireSquadHandler{
		botService: botService,
	}
}

func (handler HireSquadHandler) Handle(gameContext *models.GameContext) {
	var text string
	for _, warrior := range models.GetWarriors() {
		text = text + warrior.GetName() + ", оружие: " + warrior.GetWeapon().GetName() + " " + COMMAND_WARRIOR + strconv.Itoa(warrior.GetID()) + "\n"
	}
	handler.botService.Send(
		gameContext.GetUpdate().
			SetText(text).
			SetUpdateType(models.MESSAGE_SIMPLE))
}
