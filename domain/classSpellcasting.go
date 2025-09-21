package domain

type ClassSpellcastingInfo struct {
	maxPreparedSpells   int
	spellsKnownPerLevel [20]int
	spellSlotsPerLevel  [20][9]int
}
