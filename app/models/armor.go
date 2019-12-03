package models

type Armor struct {
	style string
	protectionValue int
}

func NewArmor(style string, protectionValue int) *Armor {
	return &Armor{
		style:   style,
		protectionValue: protectionValue,
	}
}

func (armor *Armor) GetStyle() string {
	return armor.style
}

func (armor *Armor) GetProtectionValue() int {
	return armor.protectionValue
}