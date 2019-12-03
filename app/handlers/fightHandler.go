package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"strconv"
)

const (
	COMMAND_START_FIGHT = "Да, сражаемся!"
	COMMAND_SKIP_FIGHT = "Нет, не сражаемся"
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
	for _, warrior := range warriors {
		enemyWarrior := enemyIsland.GetWarriors()[enemyIndex]
		ourWeapon := warrior.GetWeapon()
		enemyArmor := enemyWarrior.GetArmor()
		damageValue := 0
		if ourWeapon.GetStyle() == enemyArmor.GetStyle() || enemyArmor.GetStyle() == models.STYLE_UNIVERSAL {
			damageValue = enemyArmor.GetProtectionValue() - ourWeapon.GetDamageValue()
		} else {
			damageValue = ourWeapon.GetDamageValue()
		}
		if damageValue >= 0 {
			handler.botService.Send(
				gameContext.GetUpdate().
					SetText(warrior.GetName() + " не нанёс урона").
					SetUpdateType(models.MESSAGE_SIMPLE))
			} else {
				handler.botService.Send(
					gameContext.GetUpdate().
						SetText(warrior.GetName() + " нанёс " + strconv.Itoa(-damageValue) + " единицы урона").
						SetUpdateType(models.MESSAGE_SIMPLE))
		}
		enemyIndex++
	}
}