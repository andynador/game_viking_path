package models

const (
	WEAPON_NAME_AX = "секира"
	WEAPON_NAME_SPEAR = "копье"
	WEAPON_NAME_BATON = "дубина"
)

type Weapon struct {
	style string
	name string
	damangeValue int
}

func NewWeapon(style string, name string, damangeValue int) *Weapon {
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

func (weapon *Weapon) GetDamageValue() int {
	return weapon.damangeValue
}
