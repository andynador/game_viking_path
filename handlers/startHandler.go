package handlers

import (
	"github.com/andynador/game_viking_path/service"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type StartHandler struct {
	bot service.Bot
}

func NewStartHandler(bot service.Bot) StartHandler {
	return StartHandler{
		bot: bot,
	}
}

func (handler StartHandler) Handle(update tgbotapi.Update) {
	handler.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, Викинг!"))
}
