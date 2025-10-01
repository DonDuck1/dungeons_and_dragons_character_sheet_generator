package domain

type Weapon struct {
	name           string
	weaponCategory string
	normalRange    int
	twoHanded      bool
}

func NewWeapon(name string, weaponCategory string, normalRange int, twoHanded bool) Weapon {
	return Weapon{name: name, weaponCategory: weaponCategory, normalRange: normalRange, twoHanded: twoHanded}
}

func (weapon Weapon) GetNumberOfOccupiedHandSlots() int {
	if weapon.twoHanded {
		return 2
	}

	return 1
}
