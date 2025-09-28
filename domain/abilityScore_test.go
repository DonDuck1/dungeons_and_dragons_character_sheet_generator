package domain

import (
	"reflect"
	"testing"
)

func TestCreateAbilityScoreListWithImprovements(t *testing.T) {
	expected := AbilityScoreList{
		strength: AbilityScore{
			name:        STRENGTH,
			base_value:  15,
			final_value: 18,
			modifier:    4,
		},
		dexterity: AbilityScore{
			name:        DEXTERITY,
			base_value:  1,
			final_value: 1,
			modifier:    -5,
		},
		constitution: AbilityScore{
			name:        CONSTITUTION,
			base_value:  20,
			final_value: 20,
			modifier:    5,
		},
		intelligence: AbilityScore{
			name:        INTELLIGENCE,
			base_value:  12,
			final_value: 12,
			modifier:    1,
		},
		wisdom: AbilityScore{
			name:        WISDOM,
			base_value:  10,
			final_value: 15,
			modifier:    2,
		},
		charisma: AbilityScore{
			name:        CHARISMA,
			base_value:  8,
			final_value: 7,
			modifier:    -2,
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
		t.Errorf("Expected %+v, got %+v", expected, abilityScoreList)
	}
}
