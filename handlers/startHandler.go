package handlers

import (
	"github.com/andynador/game_viking_path/service/bot"
)

type StartHandler struct {
}

func NewStartHandler() *StartHandler {
	return &StartHandler{}
}

func (handler StartHandler) Handle(update bot.Update) {
	bot.Send(update)
}
