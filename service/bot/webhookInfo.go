package bot

import "github.com/go-telegram-bot-api/telegram-bot-api"

type WebhookInfo struct {
	parent tgbotapi.WebhookInfo
}

func (webhookInfo WebhookInfo) GetLastErrorDate() int {
	return webhookInfo.parent.LastErrorDate
}
