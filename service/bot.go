package service

import "github.com/go-telegram-bot-api/telegram-bot-api"

var (
	Bot *tgbotapi.BotAPI
)

func Init(token string) error {
	var err error
	Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	return nil
}

