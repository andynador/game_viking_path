package interfaces

import (
	"github.com/andynador/game_viking_path/app/models"
)

type BotServiceAdapter interface {
	SetDebug(debug bool)
	GetUserName() string
	SetWebhook(url string) error
	GetWebhookInfo() (models.WebhookInfo, error)
	Send(update *models.Update) error
	Send(update *models.Update) error
}
