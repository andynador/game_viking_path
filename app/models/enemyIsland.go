package models

type EnemyIsland struct {
	id int
	warriors []*Warrior
}


func NewEnemyIsland(id int) *EnemyIsland {
	return &EnemyIsland{
		id:       id,
		warriors: make([]*Warrior, 0),
	}
}

func (enemyIsland *EnemyIsland) GetID() int {
	return enemyIsland.id
}

func (enemyIsland *EnemyIsland) AddWarrior(warrior *Warrior) *EnemyIsland {
	enemyIsland.warriors = append(enemyIsland.warriors, warrior)

	return enemyIsland
}

func (enemyIsland *EnemyIsland) GetWarriors() []*Warrior {
	return enemyIsland.warriors
}