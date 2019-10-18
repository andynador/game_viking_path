package main

import (
	"encoding/json"
	"fmt"
	"github.com/andynador/game_viking_path/handlers"
	"github.com/andynador/game_viking_path/service/bot"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := os.Getenv("BOT_TOKEN")

	err := bot.Init(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.SetDebug(true)

	log.Printf("Authorized on account %s", bot.GetUserName())

	err = bot.SetWebhook(os.Getenv("BOT_WEBHOOK_HOST") + "/" + token + "/webhook")
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.GetLastErrorDate() != 0 {
		log.Printf("Telegram callback failed: %d", info.GetLastErrorDate())
	}
	http.HandleFunc("/" + token + "/webhook", handlerWebhook)
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
		handler := handlers.NewStartHandler()
		handler.Handle(bot.NewUpdate(update.Message.Chat.ID, "Привет, Викинг!"))
	}
	fmt.Println(update.Message.Text)
}