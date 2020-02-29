package models

const (
	WEAPON_NAME_AX = "секира"
	WEAPON_NAME_SPEAR = "копье"
	WEAPON_NAME_BATON = "дубина"
)

type Weapon struct {
	style string
	name string
	damangeValue float32
}

func NewWeapon(style string, name string, damangeValue float32) *Weapon {
	return &Weapon{
		style:   style,
		name: name,
		damangeValue: damangeValue,
	}
}

func (weapon *Weapon) GetStyle() string {
	return weapon.style
}

func (weapon *Weapon) GetName() string {
	return weapon.name
}

func (weapon *Weapon) GetDamageValue() float32 {
	return weapon.damangeValue
}
