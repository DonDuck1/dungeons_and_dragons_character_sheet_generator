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
