package main

import (
	"encoding/json"
	"fmt"
	"github.com/andynador/game_viking_path/app/adapters"
	"github.com/andynador/game_viking_path/app/handlers"
	"github.com/andynador/game_viking_path/app/interfaces"
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/preferences"
	"github.com/andynador/game_viking_path/app/services"
	"github.com/andynador/game_viking_path/app/services/db"
	gc "github.com/andynador/game_viking_path/app/services/gameContext"
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
	invasionHandler  *handlers.InvasionHandler
	fightHandler     *handlers.FightHandler
	gameContext      *gc.GameContext
)

func main() {
	var err error
	err = initGameContext()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(gameContext.GetPreferences().BotToken)
	if err != nil {
		log.Fatal(err)
	}
	botService = services.NewBotService(adapters.NewTgBotApiAdapter(bot))

	botService.SetDebug(true)

	log.Printf("Authorized on account %s", botService.GetUserName())

	webhook := "/" + gameContext.GetPreferences().BotToken + "/webhook"

	err = botService.SetWebhook(gameContext.GetPreferences().BotWebhookHost + webhook)
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
	http.HandleFunc(webhook, handlerWebhook)
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
		gameContext = gameContext.SetUpdate(
			models.NewUpdate(update.Message.Chat.ID).SetText(update.Message.Text),
		).SetUser(getUser(update))
		go commandHandler.Handle(gameContext)
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

	if update.Message.Text == handlers.COMMAND_INVASION {
		if invasionHandler == nil {
			invasionHandler = handlers.NewInvasionHandler(botService)
		}
		return invasionHandler
	}

	if update.Message.Text == handlers.COMMAND_START_FIGHT {
		if fightHandler == nil {
			fightHandler = handlers.NewFightHandler(botService)
		}
		return fightHandler
	}

	return nil
}

func getUser(update tgbotapi.Update) models.User {
	user, isExists, err := models.GetUserByExternalId(gameContext.GetDB(), update.Message.From.ID)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !isExists {
		user, err = models.CreateUser(gameContext.GetDB(), update.Message.From.ID, update.Message.From.UserName)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	return user
}

func initGameContext() ( error) {
	gameContext = gc.New()
	p, err := preferences.Get()

	if err != nil {
		return err
	}

	db, err := db.New(p.DatabaseURL, db.Config{MaxConnLifetimeSec: p.DatabaseMaxConnLifetimeSec, MaxIdleConns: p.DatabaseMaxIdleConns, MaxOpenConns: p.DatabaseMaxOpenConns})

	if err != nil {
		return err
	}

	err = db.Connect()
	if err != nil {
		return err
	}

	gameContext = gameContext.SetDB(db).
		SetPreferences(p)

	return nil
}
