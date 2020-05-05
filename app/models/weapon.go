package models

type Weapon struct {
	style       string
	name        string
	damageValue float32
}

func (weapon Weapon) GetStyle() string {
	return weapon.style
}

func (weapon Weapon) GetName() string {
	return weapon.name
}

func (weapon Weapon) GetDamageValue() float32 {
	return weapon.damageValue
}
