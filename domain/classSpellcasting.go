package domain

import (
	"fmt"
	"math"
)

type ClassSpellcastingInfo struct {
	maxKnownSpells      *int
	maxPreparedSpells   *int
	spellList           SpellList
	spellSlotAmount     [9]int
	spellcastingAbility AbilityScoreName
	spellSaveDC         int
	spellAttackBonus    int
}

func NewClassSpellcastingInfo(className ClassName, classLevel int, abilityScoreList AbilityScoreList, proficiencyBonus int) (*ClassSpellcastingInfo, error) {
	if !(classLevel >= 1 && classLevel <= 20) {
		err := fmt.Errorf("Unvalid level provided: %d", classLevel)
		return nil, err
	}

	switch className {
	case BARD:
		spellcastingAbility := CHARISMA
		spellcastingAbilityModifier := abilityScoreList.charisma.modifier

		maxKnownSpells := BardSpellsKnownAtLevel(classLevel)

		return &ClassSpellcastingInfo{
			maxKnownSpells:      &maxKnownSpells,
			maxPreparedSpells:   nil,
			spellList:           NewSpellList(),
			spellSlotAmount:     FullSpellcasterSpellSlotsAtLevel(classLevel),
			spellcastingAbility: spellcastingAbility,
			spellSaveDC:         8 + proficiencyBonus + spellcastingAbilityModifier,
			spellAttackBonus:    proficiencyBonus + spellcastingAbilityModifier,
		}, nil
	case CLERIC:
		spellcastingAbility := WISDOM
		spellcastingAbilityModifier := abilityScoreList.wisdom.modifier

		maxPreparedSpells := max(1, spellcastingAbilityModifier+classLevel)

		return &ClassSpellcastingInfo{
			maxKnownSpells:      nil,
			maxPreparedSpells:   &maxPreparedSpells,
			spellList:           NewSpellList(),
			spellSlotAmount:     FullSpellcasterSpellSlotsAtLevel(classLevel),
			spellcastingAbility: spellcastingAbility,
			spellSaveDC:         8 + proficiencyBonus + spellcastingAbilityModifier,
			spellAttackBonus:    proficiencyBonus + spellcastingAbilityModifier,
		}, nil
	case DRUID:
		spellcastingAbility := WISDOM
		spellcastingAbilityModifier := abilityScoreList.wisdom.modifier

		maxPreparedSpells := max(1, spellcastingAbilityModifier+classLevel)

		return &ClassSpellcastingInfo{
			maxKnownSpells:      nil,
			maxPreparedSpells:   &maxPreparedSpells,
			spellList:           NewSpellList(),
			spellSlotAmount:     FullSpellcasterSpellSlotsAtLevel(classLevel),
			spellcastingAbility: spellcastingAbility,
			spellSaveDC:         8 + proficiencyBonus + spellcastingAbilityModifier,
			spellAttackBonus:    proficiencyBonus + spellcastingAbilityModifier,
		}, nil
	case PALADIN:
		spellcastingAbility := CHARISMA
		spellcastingAbilityModifier := abilityScoreList.charisma.modifier

		maxPreparedSpells := max(1, int(math.Floor(float64(spellcastingAbilityModifier)+(float64(classLevel)/2))))

		return &ClassSpellcastingInfo{
			maxKnownSpells:      nil,
			maxPreparedSpells:   &maxPreparedSpells,
			spellList:           NewSpellList(),
			spellSlotAmount:     PartialSpellcasterSpellSlotsAtLevel(classLevel),
			spellcastingAbility: spellcastingAbility,
			spellSaveDC:         8 + proficiencyBonus + spellcastingAbilityModifier,
			spellAttackBonus:    proficiencyBonus + spellcastingAbilityModifier,
		}, nil
	case RANGER:
		spellcastingAbility := WISDOM
		spellcastingAbilityModifier := abilityScoreList.wisdom.modifier

		maxKnownSpells := RangerSpellsKnownAtLevel(classLevel)

		return &ClassSpellcastingInfo{
			maxKnownSpells:      &maxKnownSpells,
			maxPreparedSpells:   nil,
			spellList:           NewSpellList(),
			spellSlotAmount:     PartialSpellcasterSpellSlotsAtLevel(classLevel),
			spellcastingAbility: spellcastingAbility,
			spellSaveDC:         8 + proficiencyBonus + spellcastingAbilityModifier,
			spellAttackBonus:    proficiencyBonus + spellcastingAbilityModifier,
		}, nil
	case SORCERER:
		spellcastingAbility := CHARISMA
		spellcastingAbilityModifier := abilityScoreList.charisma.modifier

		maxKnownSpells := SorcererSpellsKnownAtLevel(classLevel)

		return &ClassSpellcastingInfo{
			maxKnownSpells:      &maxKnownSpells,
			maxPreparedSpells:   nil,
			spellList:           NewSpellList(),
			spellSlotAmount:     FullSpellcasterSpellSlotsAtLevel(classLevel),
			spellcastingAbility: spellcastingAbility,
			spellSaveDC:         8 + proficiencyBonus + spellcastingAbilityModifier,
			spellAttackBonus:    proficiencyBonus + spellcastingAbilityModifier,
		}, nil
	case WIZARD:
		spellcastingAbility := INTELLIGENCE
		spellcastingAbilityModifier := abilityScoreList.intelligence.modifier

		maxPreparedSpells := max(1, spellcastingAbilityModifier+classLevel)

		return &ClassSpellcastingInfo{
			maxKnownSpells:      nil,
			maxPreparedSpells:   &maxPreparedSpells,
			spellList:           NewSpellList(),
			spellSlotAmount:     FullSpellcasterSpellSlotsAtLevel(classLevel),
			spellcastingAbility: spellcastingAbility,
			spellSaveDC:         8 + proficiencyBonus + spellcastingAbilityModifier,
			spellAttackBonus:    proficiencyBonus + spellcastingAbilityModifier,
		}, nil
	case BARBARIAN, FIGHTER, MONK, ROGUE, WARLOCK:
		return nil, nil
	default:
		err := fmt.Errorf("Unknown class provided: %s", className)
		return nil, err
	}
}
