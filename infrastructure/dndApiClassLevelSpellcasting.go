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

func NewDndApiClassLevelSpellcasting(
	cantripsKnown *int,
	spellsKnown *int,
	spellSlotsLevel1 int,
	spellSlotsLevel2 int,
	spellSlotsLevel3 int,
	spellSlotsLevel4 int,
	spellSlotsLevel5 int,
	spellSlotsLevel6 *int,
	spellSlotsLevel7 *int,
	spellSlotsLevel8 *int,
	spellSlotsLevel9 *int,
) DndApiClassLevelSpellcasting {
	return DndApiClassLevelSpellcasting{
		CantripsKnown:    cantripsKnown,
		SpellsKnown:      spellsKnown,
		SpellSlotsLevel1: spellSlotsLevel1,
		SpellSlotsLevel2: spellSlotsLevel2,
		SpellSlotsLevel3: spellSlotsLevel3,
		SpellSlotsLevel4: spellSlotsLevel4,
		SpellSlotsLevel5: spellSlotsLevel5,
		SpellSlotsLevel6: spellSlotsLevel6,
		SpellSlotsLevel7: spellSlotsLevel7,
		SpellSlotsLevel8: spellSlotsLevel8,
		SpellSlotsLevel9: spellSlotsLevel9,
	}
}

func (dndApiClassLevelSpellcasting DndApiClassLevelSpellcasting) GetDeepCopy() DndApiClassLevelSpellcasting {
	var deepCopiedCantripsKnown *int
	if dndApiClassLevelSpellcasting.CantripsKnown != nil {
		value := *dndApiClassLevelSpellcasting.CantripsKnown
		deepCopiedCantripsKnown = &value
	}

	var deepCopiedSpellsKnown *int
	if dndApiClassLevelSpellcasting.SpellsKnown != nil {
		value := *dndApiClassLevelSpellcasting.SpellsKnown
		deepCopiedSpellsKnown = &value
	}

	var deepCopiedSpellSlotsLevel6 *int
	if dndApiClassLevelSpellcasting.SpellSlotsLevel6 != nil {
		value := *dndApiClassLevelSpellcasting.SpellSlotsLevel6
		deepCopiedSpellSlotsLevel6 = &value
	}

	var deepCopiedSpellSlotsLevel7 *int
	if dndApiClassLevelSpellcasting.SpellSlotsLevel7 != nil {
		value := *dndApiClassLevelSpellcasting.SpellSlotsLevel7
		deepCopiedSpellSlotsLevel7 = &value
	}

	var deepCopiedSpellSlotsLevel8 *int
	if dndApiClassLevelSpellcasting.SpellSlotsLevel8 != nil {
		value := *dndApiClassLevelSpellcasting.SpellSlotsLevel8
		deepCopiedSpellSlotsLevel8 = &value
	}

	var deepCopiedSpellSlotsLevel9 *int
	if dndApiClassLevelSpellcasting.SpellSlotsLevel9 != nil {
		value := *dndApiClassLevelSpellcasting.SpellSlotsLevel9
		deepCopiedSpellSlotsLevel9 = &value
	}

	return NewDndApiClassLevelSpellcasting(
		deepCopiedCantripsKnown,
		deepCopiedSpellsKnown,
		dndApiClassLevelSpellcasting.SpellSlotsLevel1,
		dndApiClassLevelSpellcasting.SpellSlotsLevel2,
		dndApiClassLevelSpellcasting.SpellSlotsLevel3,
		dndApiClassLevelSpellcasting.SpellSlotsLevel4,
		dndApiClassLevelSpellcasting.SpellSlotsLevel5,
		deepCopiedSpellSlotsLevel6,
		deepCopiedSpellSlotsLevel7,
		deepCopiedSpellSlotsLevel8,
		deepCopiedSpellSlotsLevel9,
	)
}
