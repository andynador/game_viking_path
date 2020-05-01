package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
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

func (handler InvasionHandler) Handle(gameContext *models.GameContext) {
	var text string
	warriors := gameContext.GetUser().GetWarriors()
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
	enemyIsland := newEnemyIsland(gameContext.GetUser())
	for _, warrior := range enemyIsland.GetWarriors() {
		text = text + warrior.GetName() + ", оружие: " + warrior.GetWeapon().GetName() + "\n"
	}
	gameContext = gameContext.SetEnemyIsland(enemyIsland)
	handler.botService.Send(
		gameContext.GetUpdate().
			SetText("Доплыли, здесь сидят \n" + text + "\n Сражаемся?").
			SetUpdateType(models.MESSAGE_WITH_KEYBOARD).
			AddKeyboardRows(models.NewKeyboardButtonRow(
				models.NewKeyboardButton(COMMAND_START_FIGHT),
				models.NewKeyboardButton(COMMAND_SKIP_FIGHT),
			)))
}

func newEnemyIsland(user *models.User) *models.EnemyIsland {
	enemyIsland := models.NewEnemyIsland(1)
	allWarriors := models.GetWarriors()
	userWarriors := user.GetWarriors()

	for _, allWarrior := range allWarriors {
		for _, userWarrior := range userWarriors {
			if allWarrior.GetID() == userWarrior.GetID() {
				continue
			}
			return enemyIsland.AddWarrior(allWarrior)
		}
	}

	return enemyIsland
}
