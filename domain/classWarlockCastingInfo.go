package domain

type ClassWarlockCastingInfo struct {
	MaxKnownCantrips    int
	MaxKnownSpells      int
	SpellList           SpellList
	SpellSlotAmount     int
	SpellSlotLevel      int
	SpellcastingAbility *AbilityScore
	SpellSaveDC         int
	SpellAttackBonus    int
}

func NewClassWarlockCastingInfo(maxKnownCantrips int, maxKnownSpells int, spellList SpellList, spellSlotAmount int, spellSlotLevel int, spellcastingAbility *AbilityScore, spellSaveDC int, spellAttackBonus int) ClassWarlockCastingInfo {
	return ClassWarlockCastingInfo{
		MaxKnownCantrips:    maxKnownCantrips,
		MaxKnownSpells:      maxKnownSpells,
		SpellList:           spellList,
		SpellSlotAmount:     spellSlotAmount,
		SpellSlotLevel:      spellSlotLevel,
		SpellcastingAbility: spellcastingAbility,
		SpellSaveDC:         spellSaveDC,
		SpellAttackBonus:    spellAttackBonus,
	}
}
