package main

import (
	"encoding/json"
	"fmt"
	"github.com/andynador/game_viking_path/app/adapters"
	"github.com/andynador/game_viking_path/app/handlers"
	"github.com/andynador/game_viking_path/app/interfaces"
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/services"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	botService       *services.BotService
	startHandler     *handlers.StartHandler
	islandHandler    *handlers.IslandHandler
	viewSquadHandler *handlers.ViewSquadHandler
	hireSquadHandler *handlers.HireSquadHandler
	warriorHandler   *handlers.WarriorHandler
	users            map[int]*models.User
)

func main() {
	var err error
	users = make(map[int]*models.User, 0)
	models.InitWarriors()
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

	commandHandler := getCommandHandler(update)
	if commandHandler != nil {
		commandHandler.Handle(models.NewUpdate(update.Message.Chat.ID).SetText(update.Message.Text), getUser(update))
	}

	fmt.Println(fmt.Sprintf("%+v", update.Message.From))
}

func getCommandHandler(update tgbotapi.Update) interfaces.HandlerInterface {
	if update.Message.Text == handlers.COMMAND_START {
		if startHandler == nil {
			startHandler = handlers.NewStartHandler(botService)
		}

		return startHandler
	}

	if update.Message.Text == handlers.COMMAND_ISLAND {
		if islandHandler == nil {
			islandHandler = handlers.NewIslandHandler(botService)
		}

		return islandHandler
	}

	if update.Message.Text == handlers.COMMAND_VIEW_SQUAD {
		if viewSquadHandler == nil {
			viewSquadHandler = handlers.NewViewSquadHandler(botService)
		}

		return viewSquadHandler
	}

	if update.Message.Text == handlers.COMMAND_HIRE_SQUAD {
		if hireSquadHandler == nil {
			hireSquadHandler = handlers.NewHireSquadHandler(botService)
		}
		return hireSquadHandler
	}

	if strings.Index(update.Message.Text, handlers.COMMAND_WARRIOR) == 0 {
		if warriorHandler == nil {
			warriorHandler = handlers.NewWarriorHandler(botService)
		}
		return warriorHandler
	}

	return nil
}

func getUser(update tgbotapi.Update) *models.User {
	if user, ok := users[update.Message.From.ID]; ok {
		return user
	}
	users[update.Message.From.ID] = models.NewUser(update.Message.From.ID, update.Message.From.UserName)

	return users[update.Message.From.ID]
}
