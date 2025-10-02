package domain

type Shield struct {
	Name               string
	ArmorClassModifier int
}

func NewShield(name string, armorClassModifier int) Shield {
	return Shield{Name: name, ArmorClassModifier: armorClassModifier}
}

func (shield Shield) GetNumberOfOccupiedHandSlots() int {
	return 1
}
