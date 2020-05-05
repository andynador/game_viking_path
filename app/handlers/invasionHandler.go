package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"github.com/andynador/game_viking_path/app/services/gameContext"
	"log"
	"time"
)

const COMMAND_INVASION = "Идём в набег"

type InvasionHandler struct {
	botService *services.BotService
}

func NewInvasionHandler(botService *services.BotService) *InvasionHandler {
	return &InvasionHandler{
		botService: botService,
	}
}

func (handler InvasionHandler) Handle(gameContext *gameContext.GameContext) {
	var text string
	warriors, err := models.GetWarriorsByUserId(gameContext.GetDB(), gameContext.GetUser().GetId())
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(warriors) == 0 {
		handler.botService.Send(
			gameContext.GetUpdate().
				SetText("У тебя пока нет ни одного война").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	handler.botService.Send(
		gameContext.GetUpdate().
			SetText("Плывём 10 секунд...").
			SetUpdateType(models.MESSAGE_SIMPLE))
	time.Sleep(10 * time.Second)
	enemyIsland, _, err := models.GetEnemyIsland(gameContext.GetDB())
	if err != nil {
		log.Fatal(err)
		return
	}
	warriors, err = models.GetWarriorsByEnemyIslandId(gameContext.GetDB(), enemyIsland.GetId())
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, warrior := range warriors {
		text = text + warrior.GetName() + ", оружие: " + warrior.GetWeapon().GetName() + "\n"
	}
	gameContext = gameContext.SetEnemyIsland(&enemyIsland)
	handler.botService.Send(
		gameContext.GetUpdate().
			SetText("Доплыли, здесь сидят \n" + text + "\n Сражаемся?").
			SetUpdateType(models.MESSAGE_WITH_KEYBOARD).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton(COMMAND_START_FIGHT),
				models.NewKeyboardButton(COMMAND_SKIP_FIGHT),
			)))
}