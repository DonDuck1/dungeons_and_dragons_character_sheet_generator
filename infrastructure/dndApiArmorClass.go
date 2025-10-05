package infrastructure

type DndApiArmorClass struct {
	Base     int  `json:"base"`
	DexBonus bool `json:"dex_bonus"`
	MaxBonus *int `json:"max_bonus"`
}
