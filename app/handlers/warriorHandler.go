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

func (handler WarriorHandler) Handle(gameContext *models.GameContext) {
	warriors := gameContext.GetUser().GetWarriors()
	if len(warriors) == 1 {
		handler.botService.Send(
			gameContext.GetUpdate().
				SetText("У тебя уже есть один воин, больше нельзя").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	text := strings.Replace(gameContext.GetUpdate().GetText(), COMMAND_WARRIOR, "", 1)
	id, _ := strconv.Atoi(text)
	warrior := models.GetWarrior(id)
	if warrior == nil {
		handler.botService.Send(
			gameContext.GetUpdate().
				SetText("Этот воин нам неизвестен").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	gameContext.GetUser().AddWarrior(warrior)

	handler.botService.Send(
		gameContext.GetUpdate().
			SetText("Воин " + warrior.GetName() + " нанят!").
			SetUpdateType(models.MESSAGE_SIMPLE))
}
