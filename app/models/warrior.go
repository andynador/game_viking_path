package models

var warriors map[int]*Warrior

type Warrior struct {
	id   int
	name string
}

func NewWarrior(id int, name string) *Warrior {
	return &Warrior{
		id:   id,
		name: name,
	}
}

func (warrior *Warrior) GetID() int {
	return warrior.id
}

func (warrior *Warrior) GetName() string {
	return warrior.name
}

func InitWarriors() {
	warriors = make(map[int]*Warrior, 0)

	warriors[1] = NewWarrior(1, "Харольд Большая секира")
	warriors[2] = NewWarrior(2, "Олав Рыжая борода")
	warriors[3] = NewWarrior(3, "Хакон Длинный язык")
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
