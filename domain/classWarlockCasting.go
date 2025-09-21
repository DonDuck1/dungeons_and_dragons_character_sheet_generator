package domain

type ClassWarlockCastingInfo struct {
	maxPreparedSpells   int
	spellsKnownPerLevel [20]int
	spellSlotsPerLevel  [20]int
	spellSlotLevel      int
}
