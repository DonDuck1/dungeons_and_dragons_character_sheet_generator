package domain

type ClassWarlockCastingInfo struct {
	maxKnownSpells      int
	spellList           SpellList
	spellSlotAmount     int
	spellSlotLevel      int
	spellcastingAbility AbilityScoreName
	spellSaveDC         int
	spellAttackBonus    int
}

func NewClassWarlockCastingInfo(maxKnownSpells int, spellList SpellList, spellSlotAmount int, spellSlotLevel int, spellcastingAbility AbilityScoreName, spellSaveDC int, spellAttackBonus int) ClassWarlockCastingInfo {
	return ClassWarlockCastingInfo{
		maxKnownSpells:      maxKnownSpells,
		spellList:           spellList,
		spellSlotAmount:     spellSlotAmount,
		spellSlotLevel:      spellSlotLevel,
		spellcastingAbility: spellcastingAbility,
		spellSaveDC:         spellSaveDC,
		spellAttackBonus:    spellAttackBonus,
	}
}

// func NewClassWarlockCastingInfo(className ClassName, classLevel int, abilityScoreList AbilityScoreList, proficiencyBonus int) (*ClassWarlockCastingInfo, error) {
// 	if !(classLevel >= 1 && classLevel <= 20) {
// 		err := fmt.Errorf("Unvalid level provided: %d", classLevel)
// 		return nil, err
// 	}

// 	switch className {
// 	case WARLOCK:
// 		spellcastingAbility := CHARISMA
// 		spellcastingAbilityModifier := abilityScoreList.charisma.modifier

// 		maxKnownSpells := WarlockSpellsKnownAtLevel(classLevel)

// 		spellSlots := WarlockSpellSlotsAtLevel(classLevel)
// 		spellSlotAmount := 0
// 		spellSlotLevel := 0

// 		for index, value := range spellSlots {
// 			if value != 0 {
// 				spellSlotAmount = value
// 				spellSlotLevel = index + 1
// 				break
// 			}
// 		}

// 		return &ClassWarlockCastingInfo{
// 			maxKnownSpells:      maxKnownSpells,
// 			spellList:           NewSpellList(),
// 			spellSlotAmount:     spellSlotAmount,
// 			spellSlotLevel:      spellSlotLevel,
// 			spellcastingAbility: spellcastingAbility,
// 			spellSaveDC:         8 + proficiencyBonus + spellcastingAbilityModifier,
// 			spellAttackBonus:    proficiencyBonus + spellcastingAbilityModifier,
// 		}, nil
// 	case BARBARIAN, BARD, CLERIC, DRUID, FIGHTER, MONK, PALADIN, RANGER, ROGUE, SORCERER, WIZARD:
// 		return nil, nil
// 	default:
// 		err := fmt.Errorf("Unknown class provided: %s", className)
// 		return nil, err
// 	}
// }
