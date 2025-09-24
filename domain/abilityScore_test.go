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
			base_value:  14,
			final_value: 14,
			modifier:    2,
		},
		constitution: AbilityScore{
			name:        CONSTITUTION,
			base_value:  13,
			final_value: 14,
			modifier:    2,
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

	abilityScoreValueList := NewAbilityScoreValueList(15, 14, 13, 12, 10, 8)
	abilityScoreImprovements := []AbilityScoreImprovement{
		NewAbilityScoreImprovement(STRENGTH, 3),
		NewAbilityScoreImprovement(CONSTITUTION, 1),
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
