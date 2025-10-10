package domain

type SpellList struct {
	Spells []Spell
}

func NewEmptySpellList() SpellList {
	spells := []Spell{}
	return SpellList{Spells: spells}
}

func NewFilledSpellList(spells []Spell) SpellList {
	return SpellList{Spells: spells}
}

func (spellList SpellList) GetPreparedSpells() []Spell {
	preparedSpells := []Spell{}
	for _, spell := range spellList.Spells {
		if spell.Prepared {
			preparedSpells = append(preparedSpells, spell)
		}
	}

	return preparedSpells
}

func (spellList *SpellList) GetAmountOfKnownCantrips() int {
	amountOfKnownCantrips := 0

	for _, spell := range spellList.Spells {
		if spell.Level == 0 {
			amountOfKnownCantrips += 1
		}
	}

	return amountOfKnownCantrips
}

func (spellList *SpellList) GetAmountOfKnownSpells() int {
	amountOfKnownSpells := 0

	for _, spell := range spellList.Spells {
		if spell.Level != 0 {
			amountOfKnownSpells += 1
		}
	}

	return amountOfKnownSpells
}

func (spellList *SpellList) AddSpell(spell Spell) {
	spellList.Spells = append(spellList.Spells, spell)
}
