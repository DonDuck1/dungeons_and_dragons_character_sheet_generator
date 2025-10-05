package infrastructure

type DndApiClassLevelSpellcasting struct {
	CantripsKnown    *int `json:"cantrips_known"`
	SpellsKnown      *int `json:"spells_known"`
	SpellSlotsLevel1 int  `json:"spell_slots_level_1"`
	SpellSlotsLevel2 int  `json:"spell_slots_level_2"`
	SpellSlotsLevel3 int  `json:"spell_slots_level_3"`
	SpellSlotsLevel4 int  `json:"spell_slots_level_4"`
	SpellSlotsLevel5 int  `json:"spell_slots_level_5"`
	SpellSlotsLevel6 *int `json:"spell_slots_level_6"`
	SpellSlotsLevel7 *int `json:"spell_slots_level_7"`
	SpellSlotsLevel8 *int `json:"spell_slots_level_8"`
	SpellSlotsLevel9 *int `json:"spell_slots_level_9"`
}
