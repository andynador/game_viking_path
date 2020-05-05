package gameContext

import (
	"github.com/andynador/game_viking_path/app/models"
	"github.com/andynador/game_viking_path/app/preferences"
	"github.com/andynador/game_viking_path/app/services/db"
)

type GameContext struct {
	update      *models.Update
	user        models.User
	enemyIsland *models.EnemyIsland
	db          *db.Database
	preferences *preferences.Preferences
}

func New() *GameContext {
	return &GameContext{}
}

func (gameContext *GameContext) SetUpdate(update *models.Update) *GameContext {
	gameContext.update = update

	return gameContext
}

func (gameContext *GameContext) GetUpdate() *models.Update {
	return gameContext.update
}

func (gameContext *GameContext) SetUser(user models.User) *GameContext {
	gameContext.user = user

	return gameContext
}

func (gameContext *GameContext) GetUser() models.User {
	return gameContext.user
}

func (gameContext *GameContext) SetEnemyIsland(enemyIsland *models.EnemyIsland) *GameContext {
	gameContext.enemyIsland = enemyIsland

	return gameContext
}

func (gameContext *GameContext) GetEnemyIsland() *models.EnemyIsland {
	return gameContext.enemyIsland
}

func (gameContext *GameContext) SetDB(db *db.Database) *GameContext {
	gameContext.db = db

	return gameContext
}

func (gameContext *GameContext) GetDB() *db.Database {
	return gameContext.db
}

func (gameContext *GameContext) SetPreferences(preferences *preferences.Preferences) *GameContext {
	gameContext.preferences = preferences

	return gameContext
}

func (gameContext *GameContext) GetPreferences()  *preferences.Preferences {
	return gameContext.preferences
}
