package domain

import (
	"fmt"
	"strings"
)

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

func (spellList *SpellList) GetAmountOfPreparedSpells() int {
	amountOfPreparedSpells := 0

	for _, spell := range spellList.Spells {
		if spell.Level != 0 && spell.Prepared {
			amountOfPreparedSpells += 1
		}
	}

	return amountOfPreparedSpells
}

func (spellList *SpellList) GetByName(spellName string) (*Spell, error) {
	for i, spell := range spellList.Spells {
		if strings.EqualFold(spell.Name, spellName) {
			return &spellList.Spells[i], nil // Use index to point to actual object, not the temporary copy of the loop
		}
	}

	err := fmt.Errorf("spell \"%s\" not found", spellName)
	return nil, err
}

func (spellList *SpellList) AddSpell(spell Spell) {
	spellList.Spells = append(spellList.Spells, spell)
}

func (spellList *SpellList) ForgetSpell(spellName string) error {
	for i, spell := range spellList.Spells {
		if strings.EqualFold(spell.Name, spellName) {
			spellList.Spells = append(spellList.Spells[:i], spellList.Spells[i+1:]...)
			return nil
		}
	}

	err := fmt.Errorf("spell \"%s\" not found", spellName)
	return err
}
