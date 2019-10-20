package handlers

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
)

type StartHandler struct {
	botService *services.BotService
}

func NewStartHandler(botService *services.BotService) *StartHandler {
	return &StartHandler{
		botService: botService,
	}
}

func (handler StartHandler) Handle(update models.Update) {
	str := "Привет, Викинг!"
	handler.botService.Send(models.NewUpdate(update.GetChatID(), &str))
}
