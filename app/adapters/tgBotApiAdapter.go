package adapters

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgBotApiAdapter struct {
	tgBotApi *tgbotapi.BotAPI
}

func NewTgBotApiAdapter(tgBotApi *tgbotapi.BotAPI) TgBotApiAdapter {
	return TgBotApiAdapter{
		tgBotApi: tgBotApi,
	}
}

func (adapter TgBotApiAdapter) SetDebug(debug bool) {
	adapter.tgBotApi.Debug = debug
}

func (adapter TgBotApiAdapter) GetUserName() string {
	return adapter.tgBotApi.Self.UserName
}

func (adapter TgBotApiAdapter) SetWebhook(url string) error {
	_, err := adapter.tgBotApi.SetWebhook(tgbotapi.NewWebhook(url))
	return err
}

func (adapter TgBotApiAdapter) GetWebhookInfo() (models.WebhookInfo, error) {
	info, err := adapter.tgBotApi.GetWebhookInfo()

	return models.NewWebhookInfo(info.LastErrorDate), err
}

func (adapter TgBotApiAdapter) Send(update *models.Update) error {
	msg := tgbotapi.NewMessage(update.GetChatID(), update.GetText())

	if update.IsUpdateType(models.MESSAGE_WITH_KEYBOARD) {
		msg.ReplyMarkup = adapter.getReplyKeyboard(update)
	}

	_ , err := adapter.tgBotApi.Send(msg)

	return err
}


func (adapter TgBotApiAdapter) getReplyKeyboard(update *models.Update) tgbotapi.ReplyKeyboardMarkup {
	keyboardRows := update.GetKeyboardRows()
	keyboardButtonRows := make([][]tgbotapi.KeyboardButton, len(keyboardRows))
	for i, row := range keyboardRows {
		keyboardButtonRows[i] = make([]tgbotapi.KeyboardButton, len(row))
		keyboardButtonRow := make([]tgbotapi.KeyboardButton, len(row))
		for _, button := range row {
			keyboardButtonRow = append(keyboardButtonRow, tgbotapi.NewKeyboardButton(button.GetText()))
		}
		keyboardButtonRows[i] = tgbotapi.NewKeyboardButtonRow(keyboardButtonRow...)
	}

	return tgbotapi.NewReplyKeyboard(keyboardButtonRows...)
}