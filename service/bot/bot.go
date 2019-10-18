package bot

import "github.com/go-telegram-bot-api/telegram-bot-api"

var (
	bot *tgbotapi.BotAPI
)

func Init(token string) error {
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	return nil
}

func SetDebug(debug bool) {
	bot.Debug = debug
}

func GetUserName() string {
	return bot.Self.UserName
}

func SetWebhook(url string) error {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(url))
	return err
}

func GetWebhookInfo() (WebhookInfo, error) {
	info, err := bot.GetWebhookInfo()

	return WebhookInfo{
		parent: info,
	}, err
}


func Send(update Update) error {
	_, err := bot.Send(tgbotapi.NewMessage(update.chatID, update.text))

	return err
}
