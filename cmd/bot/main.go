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
	botService       *services.BotService
	startHandler     *handlers.StartHandler
	islandHandler    *handlers.IslandHandler
	viewSquadHandler *handlers.ViewSquadHandler
	hireSquadHandler *handlers.HireSquadHandler
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
	if update.Message.Text == handlers.COMMAND_START {
		if startHandler == nil {
			startHandler = handlers.NewStartHandler(botService)
		}
		startHandler.Handle(models.NewUpdate(update.Message.Chat.ID))
	}
	if update.Message.Text == handlers.COMMAND_ISLAND {
		if islandHandler == nil {
			islandHandler = handlers.NewIslandHandler(botService)
		}
		islandHandler.Handle(models.NewUpdate(update.Message.Chat.ID))
	}
	if update.Message.Text == handlers.COMMAND_VIEW_SQUAD {
		if viewSquadHandler == nil {
			viewSquadHandler = handlers.NewViewSquadHandler(botService)
		}
		viewSquadHandler.Handle(models.NewUpdate(update.Message.Chat.ID))
	}
	if update.Message.Text == handlers.COMMAND_HIRE_SQUAD {
		if hireSquadHandler == nil {
			hireSquadHandler = handlers.NewHireSquadHandler(botService)
		}
		hireSquadHandler.Handle(models.NewUpdate(update.Message.Chat.ID))
	}

	fmt.Println(update.Message.Text)
}
