package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type StartHandler struct {
	bot *tgbotapi.BotAPI
}

func NewStartHandler(bot *tgbotapi.BotAPI) *StartHandler {
	return &StartHandler{
		bot: bot,
	}
}

func (handler StartHandler) Handle(update tgbotapi.Update) {
	handler.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, Викинг!"))
}
