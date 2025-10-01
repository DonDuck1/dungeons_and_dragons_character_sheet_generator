package domain

type Shield struct {
	name               string
	armorClassModifier int
}

func NewShield(name string, armorClassModifier int) Shield {
	return Shield{name: name, armorClassModifier: armorClassModifier}
}

func (shield Shield) GetNumberOfOccupiedHandSlots() int {
	return 1
}
