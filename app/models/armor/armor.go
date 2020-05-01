package armor

type Armor struct {
	style           string
	protectionValue float32
}

func New(style string, protectionValue float32) *Armor {
	return &Armor{
		style:           style,
		protectionValue: protectionValue,
	}
}

func (armor *Armor) GetStyle() string {
	return armor.style
}

func (armor *Armor) GetProtectionValue() float32 {
	return armor.protectionValue
}
