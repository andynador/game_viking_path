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

func (handler InvasionHandler) Handle(update *models.Update, user *models.User) {
	warriors := user.GetWarriors()
	if len(warriors) == 0 {
		handler.botService.Send(
			update.
				SetText("У тебя пока нет ни одного война").
				SetUpdateType(models.MESSAGE_SIMPLE))
		return
	}
	time.Sleep(10 * time.Second)
	handler.botService.Send(
			update.
				SetText("Подождали 10 секунд").
				SetUpdateType(models.MESSAGE_SIMPLE))
}