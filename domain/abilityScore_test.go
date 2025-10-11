package domain

import (
	"reflect"
	"testing"
)

func TestCreateAbilityScoreListWithImprovements(t *testing.T) {
	expected := AbilityScoreList{
		Strength: AbilityScore{
			Name:       STRENGTH,
			BaseValue:  15,
			FinalValue: 18,
			Modifier:   4,
		},
		Dexterity: AbilityScore{
			Name:       DEXTERITY,
			BaseValue:  1,
			FinalValue: 1,
			Modifier:   -5,
		},
		Constitution: AbilityScore{
			Name:       CONSTITUTION,
			BaseValue:  20,
			FinalValue: 20,
			Modifier:   5,
		},
		Intelligence: AbilityScore{
			Name:       INTELLIGENCE,
			BaseValue:  12,
			FinalValue: 12,
			Modifier:   1,
		},
		Wisdom: AbilityScore{
			Name:       WISDOM,
			BaseValue:  10,
			FinalValue: 15,
			Modifier:   2,
		},
		Charisma: AbilityScore{
			Name:       CHARISMA,
			BaseValue:  8,
			FinalValue: 7,
			Modifier:   -2,
		},
	}

	abilityScoreValueList := NewAbilityScoreValueList(15, -6, 25, 12, 10, 8)
	abilityScoreImprovements := []AbilityScoreImprovement{
		NewAbilityScoreImprovement(STRENGTH, 3),
		NewAbilityScoreImprovement(DEXTERITY, -4),
		NewAbilityScoreImprovement(CONSTITUTION, 4),
		NewAbilityScoreImprovement(WISDOM, 2),
		NewAbilityScoreImprovement(WISDOM, 3),
		NewAbilityScoreImprovement(CHARISMA, -1),
	}
	abilityScoreImprovementList := NewAbilityScoreImprovementList(abilityScoreImprovements)
	abilityScoreList := NewAbilityScoreList(abilityScoreValueList, abilityScoreImprovementList)

	if !reflect.DeepEqual(expected, abilityScoreList) {
		t.Errorf("expected %+v, got %+v", expected, abilityScoreList)
	}
}
