package main

import (
	"encoding/json"
	"fmt"
	"github.com/andynador/game_viking_path/app/adapters"
	"github.com/andynador/game_viking_path/app/handlers"
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	botService   *services.BotService
	startHandler *handlers.StartHandler
)

func main() {
	var err error
	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	botService = services.NewBotService(adapters.NewTgBotApiAdapter(bot))

	botService.SetDebug(true)

	log.Printf("Authorized on account %s", botService.GetUserName())

	err = botService.SetWebhook(os.Getenv("BOT_WEBHOOK_HOST") + "/" + token + "/webhook")
	if err != nil {
		log.Fatal(err)
	}
	info, err := botService.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.GetLastErrorDate() != 0 {
		log.Printf("Telegram callback failed: %d", info.GetLastErrorDate())
	}
	http.HandleFunc("/"+token+"/webhook", handlerWebhook)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handlerWebhook(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var update tgbotapi.Update
	json.Unmarshal(bytes, &update)
	if update.Message.Text == "/start" {
		if startHandler == nil {
			startHandler = handlers.NewStartHandler(botService)
		}
		startHandler.Handle(models.NewUpdate(update.Message.Chat.ID, nil))
	}
	fmt.Println(update.Message.Text)
}
