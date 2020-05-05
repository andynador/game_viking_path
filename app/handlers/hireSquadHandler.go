package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"github.com/andynador/game_viking_path/app/services/gameContext"
	"log"
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

func (handler HireSquadHandler) Handle(gameContext *gameContext.GameContext) {
	var text string
	warriors, err := models.GetFreeWarriors(gameContext.GetDB())
	if err != nil {
		log.Fatal(err)
	}
	for _, warrior := range warriors {
		text = text + warrior.GetName() + ", оружие: " + warrior.GetWeapon().GetName() + " " + COMMAND_WARRIOR + strconv.Itoa(warrior.GetId()) + "\n"
	}
	handler.botService.Send(
		gameContext.GetUpdate().
			SetText(text).
			SetUpdateType(models.MESSAGE_SIMPLE))
}
