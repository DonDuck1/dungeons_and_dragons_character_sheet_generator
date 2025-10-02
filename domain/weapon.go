package domain

type Weapon struct {
	Name           string
	WeaponCategory string
	NormalRange    int
	TwoHanded      bool
}

func NewWeapon(name string, weaponCategory string, normalRange int, twoHanded bool) Weapon {
	return Weapon{Name: name, WeaponCategory: weaponCategory, NormalRange: normalRange, TwoHanded: twoHanded}
}

func (weapon Weapon) GetNumberOfOccupiedHandSlots() int {
	if weapon.TwoHanded {
		return 2
	}

	return 1
}
