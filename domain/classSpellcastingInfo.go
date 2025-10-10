package domain

type ClassSpellcastingInfo struct {
	MaxKnownCantrips    int
	MaxKnownSpells      *int
	MaxPreparedSpells   *int
	SpellList           SpellList
	SpellSlotAmount     [9]int
	SpellcastingAbility *AbilityScore
	SpellSaveDC         int
	SpellAttackBonus    int
}

func NewClassSpellcastingInfo(maxKnownCantrips int, maxKnownSpells *int, maxPreparedSpells *int, spellList SpellList, spellSlotAmount [9]int, spellcastingAbility *AbilityScore, spellSaveDC int, spellAttackBonus int) ClassSpellcastingInfo {
	return ClassSpellcastingInfo{
		MaxKnownCantrips:    maxKnownCantrips,
		MaxKnownSpells:      maxKnownSpells,
		MaxPreparedSpells:   maxPreparedSpells,
		SpellList:           spellList,
		SpellSlotAmount:     spellSlotAmount,
		SpellcastingAbility: spellcastingAbility,
		SpellSaveDC:         spellSaveDC,
		SpellAttackBonus:    spellAttackBonus,
	}
}

func (classSpellcastingInfo ClassSpellcastingInfo) GetHighestSpellSlotLevel() int {
	for i, spellSlotAmount := range classSpellcastingInfo.SpellSlotAmount {
		if spellSlotAmount == 0 {
			return i
		}
	}

	return 9
}
