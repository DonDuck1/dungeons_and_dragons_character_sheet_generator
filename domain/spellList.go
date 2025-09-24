package domain

type SpellList struct {
	spells []Spell
}

func NewSpellList() SpellList {
	return SpellList{spells: []Spell{}}
}
