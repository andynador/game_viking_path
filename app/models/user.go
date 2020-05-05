package models

import (
	"database/sql"
	"github.com/andynador/game_viking_path/app/services/db"
)

type User struct {
	id         int
	externalId int
	login      string
}

func (user User) GetId() int {
	return user.id
}

func (user User) GetExternalId() int {
	return user.externalId
}

func (user User) GeLogin() string {
	return user.login
}

func GetUserByExternalId(db *db.Database, externalId int) (User, bool, error) {
	var (
		user User
	)
	err := db.GetConnection().QueryRow(`select users.id, users.external_id, users.login from users where external_id = $1`, externalId).Scan(&user.id, &user.externalId, &user.login)

	if err == sql.ErrNoRows {
		return user, false, nil
	}

	if err != nil {
		return user, false, err
	}

	return user, true, nil
}

func CreateUser(db *db.Database, externalId int, login string) (User, error) {
	user := User{externalId: externalId, login: login}
	err := db.GetConnection().QueryRow("insert into users(external_id, login) values($1, $2) returning id", externalId, login).Scan(&user.id)
	if err != nil && err != sql.ErrNoRows {
		return user, err
	}

	return user, nil
}
