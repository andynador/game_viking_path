package main

import (
	"encoding/json"
	"fmt"
	"github.com/andynador/game_viking_path/handlers"
	"github.com/andynador/game_viking_path/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	err := service.Init(os.Getenv("BOT_TOKEN"))

	service.Bot.Debug = true

	log.Printf("Authorized on account %s", service.Bot.Self.UserName)

	_, err = service.Bot.SetWebhook(tgbotapi.NewWebhook(os.Getenv("BOT_WEBHOOK_HOST") + "/" + service.Bot.Token + "/webhook"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := service.Bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	http.HandleFunc("/" + service.Bot.Token + "/webhook", handlerWebhook)
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
		handler := handlers.NewStartHandler(service.Bot)
		handler.Handle(update)
	}
	fmt.Println(update.Message.Text)
}