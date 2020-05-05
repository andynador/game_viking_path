package models

import (
	"database/sql"
	"github.com/andynador/game_viking_path/app/services/db"
)

const (
	STYLE_UNIVERSAL = "universal"
)

type Warrior struct {
	id     int
	name   string
	health float32
	weapon Weapon
	armor  Armor
}

func (warrior Warrior) GetId() int {
	return warrior.id
}

func (warrior Warrior) GetName() string {
	return warrior.name
}

func (warrior Warrior) GetHealth() float32 {
	return warrior.health
}

func (warrior Warrior) GetWeapon() Weapon {
	return warrior.weapon
}

func (warrior Warrior) GetArmor() Armor {
	return warrior.armor
}

func (warrior Warrior) SubHealth(health float32) {
	warrior.health -= health
}

func (warrior Warrior) IsLive() bool {
	return warrior.health > 0
}

func GetWarriorById(db *db.Database, id int) (Warrior, bool, error) {
	var (
		warrior Warrior
	)
	err := db.GetConnection().QueryRow(`
		select warrior.id, 
			warrior.name, 
			warrior.health_value, 
			weapon.name, 
			weapon.style, 
			weapon.damage_value,
			armor.style,
			armor.protection_value
		from warrior
		inner join weapon
			on weapon.id = warrior.weapon_id
		inner join armor
			on armor.id = warrior.armor_id
		where warrior.id = $1`, id).Scan(&warrior.id, &warrior.name, &warrior.health, &warrior.weapon.name, &warrior.weapon.style, &warrior.weapon.damageValue, &warrior.armor.style, &warrior.armor.protectionValue)

	if err == sql.ErrNoRows {
		return warrior, false, nil
	}

	if err != nil {
		return warrior, false, err
	}

	return warrior, true, nil
}

func LinkWarriorToUser(db *db.Database, warriorId, userId int) error {
	_, err := db.GetConnection().Exec(`update warrior set user_id = $1 where id = $2`, &userId, &warriorId)

	return err
}

func GetFreeWarriors(db *db.Database) ([]Warrior, error) {
	var (
		warriors []Warrior
		warrior Warrior
	)
	items, err := db.GetConnection().Query(`
		select warrior.id, 
			warrior.name, 
			warrior.health_value, 
			weapon.name, 
			weapon.style, 
			weapon.damage_value,
			armor.style,
			armor.protection_value
		from warrior
		inner join weapon
			on weapon.id = warrior.weapon_id
		inner join armor
			on armor.id = warrior.armor_id
		where warrior.user_id is null and warrior.enemy_island_id is null`)

	if err != nil {
		return warriors, err
	}
	defer items.Close()

	for items.Next() {
		err = items.Scan(&warrior.id, &warrior.name, &warrior.health, &warrior.weapon.name, &warrior.weapon.style, &warrior.weapon.damageValue, &warrior.armor.style, &warrior.armor.protectionValue)
		if err != nil {
			return warriors, err
		}
		warriors = append(warriors, warrior)
	}

	return warriors, nil
}

func GetWarriorsByUserId(db *db.Database, userId int) ([]Warrior, error) {
	var (
		warriors []Warrior
		warrior Warrior
	)
	items, err := db.GetConnection().Query(`
		select warrior.id, 
			warrior.name, 
			warrior.health_value, 
			weapon.name, 
			weapon.style, 
			weapon.damage_value,
			armor.style,
			armor.protection_value
		from warrior
		inner join weapon
			on weapon.id = warrior.weapon_id
		inner join armor
			on armor.id = warrior.armor_id
		where warrior.user_id = $1`, userId)

	if err != nil {
		return warriors, err
	}
	defer items.Close()

	for items.Next() {
		err = items.Scan(&warrior.id, &warrior.name, &warrior.health, &warrior.weapon.name, &warrior.weapon.style, &warrior.weapon.damageValue, &warrior.armor.style, &warrior.armor.protectionValue)
		if err != nil {
			return warriors, err
		}
		warriors = append(warriors, warrior)
	}

	return warriors, nil
}

func GetWarriorsByEnemyIslandId(db *db.Database, enemyIslandId int) ([]Warrior, error) {
	var (
		warriors []Warrior
		warrior Warrior
	)
	items, err := db.GetConnection().Query(`
		select warrior.id, 
			warrior.name, 
			warrior.health_value, 
			weapon.name, 
			weapon.style, 
			weapon.damage_value,
			armor.style,
			armor.protection_value
		from warrior
		inner join weapon
			on weapon.id = warrior.weapon_id
		inner join armor
			on armor.id = warrior.armor_id
		where warrior.enemy_island_id = $1`, enemyIslandId)

	if err != nil {
		return warriors, err
	}
	defer items.Close()

	for items.Next() {
		err = items.Scan(&warrior.id, &warrior.name, &warrior.health, &warrior.weapon.name, &warrior.weapon.style, &warrior.weapon.damageValue, &warrior.armor.style, &warrior.armor.protectionValue)
		if err != nil {
			return warriors, err
		}
		warriors = append(warriors, warrior)
	}

	return warriors, nil
}