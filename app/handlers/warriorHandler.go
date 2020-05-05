package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"github.com/andynador/game_viking_path/app/services/gameContext"
	"log"
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

func (handler WarriorHandler) Handle(gameContext *gameContext.GameContext) {
	warriors, err := models.GetWarriorsByUserId(gameContext.GetDB(), gameContext.GetUser().GetId())
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(warriors) == 1 {
		handler.botService.Send(
			gameContext.GetUpdate().
				SetText("У тебя уже есть один воин, больше нельзя").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	text := strings.Replace(gameContext.GetUpdate().GetText(), COMMAND_WARRIOR, "", 1)
	id, _ := strconv.Atoi(text)
	warrior, isExists, err := models.GetWarriorById(gameContext.GetDB(), id)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !isExists {
		handler.botService.Send(
			gameContext.GetUpdate().
				SetText("Этот воин нам неизвестен").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	models.LinkWarriorToUser(gameContext.GetDB(), warrior.GetId(), gameContext.GetUser().GetId())
	if err != nil {
		log.Fatal(err)
		return
	}

	handler.botService.Send(
		gameContext.GetUpdate().
			SetText("Воин " + warrior.GetName() + " нанят!").
			SetUpdateType(models.MESSAGE_SIMPLE))
}
