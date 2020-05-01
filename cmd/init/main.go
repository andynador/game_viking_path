package main

import (
	"database/sql"
	"fmt"
	"github.com/andynador/game_viking_path/app/preferences"
	"github.com/andynador/game_viking_path/app/services/db"
	"os"
)

func main()  {
	var (
		id int
	)
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
	var values []interface{}
	values = append(values, "Харольд Большая секира", 50)

	err = db.GetConnection().QueryRow("INSERT INTO warrior(name, health_value) VALUES($1, $2) returning id", values...).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка вставки в БД: %s", err.Error()))
		os.Exit(1)
	}
	fmt.Println(id)
}
