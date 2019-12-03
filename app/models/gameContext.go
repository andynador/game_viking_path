package models

type GameContext struct {
	update *Update
	user *User
	enemyIsland *EnemyIsland
}

func NewGameContext() *GameContext {
	return &GameContext{}
}

func (gameContext *GameContext) SetUpdate(update *Update) *GameContext {
	gameContext.update = update

	return gameContext
}

func (gameContext *GameContext) GetUpdate() *Update {
	return gameContext.update
}

func (gameContext *GameContext) SetUser(user *User) *GameContext {
	gameContext.user = user

	return gameContext
}

func (gameContext *GameContext) GetUser() *User {
	return gameContext.user
}

func (gameContext *GameContext) SetEnemyIsland(enemyIsland *EnemyIsland) *GameContext {
	gameContext.enemyIsland = enemyIsland

	return gameContext
}

func (gameContext *GameContext) GetEnemyIsland() *EnemyIsland {
	return gameContext.enemyIsland
}