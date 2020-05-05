package models

import (
	"database/sql"
	"github.com/andynador/game_viking_path/app/services/db"
)

type EnemyIsland struct {
	id   int
	name string
}

func (enemyIsland *EnemyIsland) GetId() int {
	return enemyIsland.id
}

func GetEnemyIsland(db *db.Database) (EnemyIsland, bool, error) {
	var (
		enemyIsland EnemyIsland
	)
	err := db.GetConnection().QueryRow(`select id, name from enemy_island limit 1`).Scan(&enemyIsland.id, &enemyIsland.name)

	if err == sql.ErrNoRows {
		return enemyIsland, false, nil
	}

	if err != nil {
		return enemyIsland, false, err
	}

	return enemyIsland, true, nil
}
