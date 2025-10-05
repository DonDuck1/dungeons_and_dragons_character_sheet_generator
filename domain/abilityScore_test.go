package domain

import (
	"reflect"
	"testing"
)

func TestCreateAbilityScoreListWithImprovements(t *testing.T) {
	expected := AbilityScoreList{
		Strength: AbilityScore{
			Name:        STRENGTH,
			Base_value:  15,
			Final_value: 18,
			Modifier:    4,
		},
		Dexterity: AbilityScore{
			Name:        DEXTERITY,
			Base_value:  1,
			Final_value: 1,
			Modifier:    -5,
		},
		Constitution: AbilityScore{
			Name:        CONSTITUTION,
			Base_value:  20,
			Final_value: 20,
			Modifier:    5,
		},
		Intelligence: AbilityScore{
			Name:        INTELLIGENCE,
			Base_value:  12,
			Final_value: 12,
			Modifier:    1,
		},
		Wisdom: AbilityScore{
			Name:        WISDOM,
			Base_value:  10,
			Final_value: 15,
			Modifier:    2,
		},
		Charisma: AbilityScore{
			Name:        CHARISMA,
			Base_value:  8,
			Final_value: 7,
			Modifier:    -2,
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
