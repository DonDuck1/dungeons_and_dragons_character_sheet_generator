package domain

type InventoryWeaponSlots struct {
	MainHand *Weapon
	OffHand  *Weapon
}

func NewEmptyInventoryWeaponSlots() InventoryWeaponSlots {
	return InventoryWeaponSlots{nil, nil}
}
