package handlers

import (
	"fmt"
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
)

const (
	COMMAND_START_FIGHT = "Да, сражаемся!"
	COMMAND_SKIP_FIGHT  = "Нет, не сражаемся"
)

type FightHandler struct {
	botService *services.BotService
}

func NewFightHandler(botService *services.BotService) *FightHandler {
	return &FightHandler{
		botService: botService,
	}
}

func (handler FightHandler) Handle(gameContext *models.GameContext) {
	warriors := gameContext.GetUser().GetWarriors()
	if len(warriors) == 0 {
		handler.botService.Send(
			gameContext.GetUpdate().
				SetText("У тебя пока нет ни одного война").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	enemyIsland := gameContext.GetEnemyIsland()
	if enemyIsland == nil {
		handler.botService.Send(
			gameContext.GetUpdate().
				SetText("Не выбран остров для нападения").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	enemyIndex := 0
	enemyWarriors := enemyIsland.GetWarriors()
	for {
		for _, warrior := range warriors {
			if enemyIndex > (len(enemyWarriors) - 1) {
				enemyIndex = 0
			}
			damageValue := attactFirstWarriorToSecondWarrior(warrior, enemyWarriors[enemyIndex])
			handler.processDamageValue(damageValue, warrior, gameContext)
			damageValue = attactFirstWarriorToSecondWarrior(enemyWarriors[enemyIndex], warrior)
			handler.processDamageValue(damageValue, enemyWarriors[enemyIndex], gameContext)

			enemyIndex++
		}
		break
	}
}

func attactFirstWarriorToSecondWarrior(firstWarrior, secondWarrior *models.Warrior) float32 {
	firstWarriorWeapon := firstWarrior.GetWeapon()
	secondWarriorArmor := secondWarrior.GetArmor()

	if firstWarriorWeapon.GetStyle() == secondWarriorArmor.GetStyle() || secondWarriorArmor.GetStyle() == models.STYLE_UNIVERSAL {
		return secondWarriorArmor.GetProtectionValue() - firstWarriorWeapon.GetDamageValue()
	}

	return firstWarriorWeapon.GetDamageValue()
}

func (handler FightHandler) processDamageValue(damageValue float32, warrior *models.Warrior, gameContext *models.GameContext) {
	if damageValue >= 0 {
		handler.sendMessage(warrior.GetName()+" не нанёс урона", gameContext)
	} else {
		handler.sendMessage(warrior.GetName()+" нанёс "+fmt.Sprintf("%.1f", damageValue)+" единицы урона", gameContext)
	}
}

func (handler FightHandler) sendMessage(message string, gameContext *models.GameContext) {
	handler.botService.Send(gameContext.GetUpdate().
		SetText(message).
		SetUpdateType(models.MESSAGE_SIMPLE))
}
