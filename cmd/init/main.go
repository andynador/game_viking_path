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
		warriorId int
		weaponId int
		armorId int
		enemyIslandId int
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

	err = db.GetConnection().QueryRow("select id from weapon where name = $1", "Секира").Scan(&weaponId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into weapon(name, style, damage_value) values($1, $2, $3) returning id", "Секира", "chopping", 5).Scan(&weaponId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}

	err = db.GetConnection().QueryRow("select id from armor where style = $1 and protection_value = $2", "universal", 2).Scan(&armorId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into armor(style, protection_value) values($1, $2) returning id", "universal", 2).Scan(&armorId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}

	err = db.GetConnection().QueryRow("select id from warrior where name = $1", "Харольд Большая секира").Scan(&warriorId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into warrior(name, health_value, weapon_id, armor_id) values($1, $2, $3, $4) returning id", "Харольд Большая секира", 50, weaponId, armorId).Scan(&warriorId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}



	err = db.GetConnection().QueryRow("select id from weapon where name = $1", "Копьё").Scan(&weaponId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into weapon(name, style, damage_value) values($1, $2, $3) returning id", "Копьё", "pricking", 3).Scan(&weaponId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}

	err = db.GetConnection().QueryRow("select id from armor where style = $1 and protection_value = $2", "chopping", 5).Scan(&armorId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into armor(style, protection_value) values($1, $2) returning id", "chopping", 5).Scan(&armorId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}

	err = db.GetConnection().QueryRow("select id from warrior where name = $1", "Олав Рыжая борода").Scan(&warriorId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into warrior(name, health_value, weapon_id, armor_id) values($1, $2, $3, $4) returning id", "Олав Рыжая борода", 40, weaponId, armorId).Scan(&warriorId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}



	err = db.GetConnection().QueryRow("select id from weapon where name = $1", "Дубина").Scan(&weaponId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into weapon(name, style, damage_value) values($1, $2, $3) returning id", "Дубина", "pricking", 3).Scan(&weaponId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}

	err = db.GetConnection().QueryRow("select id from armor where style = $1 and protection_value = $2", "chopping", 4).Scan(&armorId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into armor(style, protection_value) values($1, $2) returning id", "chopping", 4).Scan(&armorId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}

	err = db.GetConnection().QueryRow("select id from enemy_island where name = $1", "Фюн").Scan(&enemyIslandId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into enemy_island(name) values($1) returning id", "Фюн").Scan(&enemyIslandId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}

	err = db.GetConnection().QueryRow("select id from warrior where name = $1", "Хакон Длинный язык").Scan(&warriorId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(fmt.Sprintf("Ошибка получения данных из БД: %s", err.Error()))
		os.Exit(1)
	}

	if err == sql.ErrNoRows {
		err = db.GetConnection().QueryRow("insert into warrior(name, health_value, enemy_island_id, weapon_id, armor_id) values($1, $2, $3, $4, $5) returning id", "Хакон Длинный язык", 45, enemyIslandId, weaponId, armorId).Scan(&warriorId)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(fmt.Sprintf("Ошибка вставки данных в БД: %s", err.Error()))
			os.Exit(1)
		}
	}
}
