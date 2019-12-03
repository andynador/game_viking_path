package models

const (
	STYLE_CHOPPING = "chopping"
	STYLE_PRICKING = "pricking"
	STYLE_PERCUSSION = "percussion"
	STYLE_UNIVERSAL = "universal"
)

var warriors map[int]*Warrior

type Warrior struct {
	id   int
	name string
	health int
	weapon *Weapon
	armor *Armor
}

func NewWarrior(id int, name string, health int, weapon *Weapon, armor *Armor) *Warrior {
	return &Warrior{
		id:   id,
		name: name,
		health: health,
		weapon: weapon,
		armor: armor,
	}
}

func (warrior *Warrior) GetID() int {
	return warrior.id
}

func (warrior *Warrior) GetName() string {
	return warrior.name
}

func (warrior *Warrior) GetHealth() int {
	return warrior.health
}

func (warrior *Warrior) GetWeapon() *Weapon {
	return warrior.weapon
}

func (warrior *Warrior) GetArmor() *Armor {
	return warrior.armor
}

func InitWarriors() {
	warriors = make(map[int]*Warrior, 0)

	warriors[1] = NewWarrior(1, 
		"Харольд Большая секира", 
		50, 
		NewWeapon(STYLE_CHOPPING, 
			WEAPON_NAME_AX, 
			5,
		),
		NewArmor(STYLE_UNIVERSAL, 2),
	)

	warriors[2] = NewWarrior(2, 
		"Олав Рыжая борода", 
		40, 
		NewWeapon(STYLE_PRICKING, 
			WEAPON_NAME_SPEAR, 
			3,
		),
		NewArmor(STYLE_CHOPPING, 5),
	)

	warriors[3] = NewWarrior(
		3, 
		"Хакон Длинный язык", 
		45,
		NewWeapon(STYLE_PRICKING, 
			WEAPON_NAME_BATON, 
			3,
		),
		NewArmor(STYLE_CHOPPING, 4),
	)
}

func GetWarrior(id int) *Warrior {
	if warrior, ok := warriors[id]; ok {
		return warrior
	}

	return nil
}

func GetWarriors() map[int]*Warrior {
	return warriors
}
