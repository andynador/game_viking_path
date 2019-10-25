package services

import (
	"github.com/andynador/game_viking_path/app/interfaces"
	"github.com/andynador/game_viking_path/app/models"
)

type BotService struct {
	botAdapter interfaces.BotServiceAdapter
}

func NewBotService(botAdapter interfaces.BotServiceAdapter) *BotService {
	return &BotService{
		botAdapter: botAdapter,
	}
}

func (botService *BotService) SetDebug(debug bool) {
	botService.botAdapter.SetDebug(debug)
}

func (botService *BotService) GetUserName() string {
	return botService.botAdapter.GetUserName()
}

func (botService *BotService) SetWebhook(url string) error {
	return botService.botAdapter.SetWebhook(url)
}

func (botService *BotService) GetWebhookInfo() (models.WebhookInfo, error) {
	return botService.botAdapter.GetWebhookInfo()
}

func (botService *BotService) Send(update *models.Update) error {
	return botService.botAdapter.Send(update)
}
