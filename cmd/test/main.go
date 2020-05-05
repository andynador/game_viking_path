package main

import (
	"fmt"
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/preferences"
	"github.com/andynador/game_viking_path/app/services/db"
	"os"
)

func main() {
	p, err := preferences.Get()
	if err != nil {
		fmt.Println(fmt.Sprintf("Ошибка парсинга настроек: %s", err.Error()))
		os.Exit(1)
	}
	db, err := db.New(p.DatabaseURL, db.Config{MaxConnLifetimeSec: p.DatabaseMaxConnLifetimeSec, MaxIdleConns: p.DatabaseMaxIdleConns, MaxOpenConns: p.DatabaseMaxOpenConns})
	if err != nil {
		fmt.Println(fmt.Sprintf("Ошибка создания объекта БД: %s", err.Error()))
		os.Exit(1)
	}
	err = db.Connect()
	if err != nil {
		fmt.Println(fmt.Sprintf("Ошибка подключения к БД: %s", err.Error()))
		os.Exit(1)
	}
	warrior, _, err := models.GetWarriorById(db, 1)
	fmt.Println(fmt.Sprintf("%+v", warrior))
}
