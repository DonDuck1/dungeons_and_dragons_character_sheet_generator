package domain

import (
	"reflect"
	"testing"
)

func TestCreateAbilityScoreListWithImprovements(t *testing.T) {
	expected := AbilityScoreList{
		strength: AbilityScore{
			name:        Strength,
			base_value:  15,
			final_value: 18,
			modifier:    4,
		},
		dexterity: AbilityScore{
			name:        Dexterity,
			base_value:  14,
			final_value: 14,
			modifier:    2,
		},
		constitution: AbilityScore{
			name:        Constitution,
			base_value:  13,
			final_value: 14,
			modifier:    2,
		},
		intelligence: AbilityScore{
			name:        Intelligence,
			base_value:  12,
			final_value: 12,
			modifier:    1,
		},
		wisdom: AbilityScore{
			name:        Wisdom,
			base_value:  10,
			final_value: 15,
			modifier:    2,
		},
		charisma: AbilityScore{
			name:        Charisma,
			base_value:  8,
			final_value: 7,
			modifier:    -2,
		},
	}

	abilityScoreValueList := NewAbilityScoreValueList(15, 14, 13, 12, 10, 8)
	abilityScoreImprovements := []AbilityScoreImprovement{
		NewAbilityScoreImprovement(Strength, 3),
		NewAbilityScoreImprovement(Constitution, 1),
		NewAbilityScoreImprovement(Wisdom, 2),
		NewAbilityScoreImprovement(Wisdom, 3),
		NewAbilityScoreImprovement(Charisma, -1),
	}
	abilityScoreImprovementList := NewAbilityScoreImprovementList(abilityScoreImprovements)
	abilityScoreList := NewAbilityScoreList(abilityScoreValueList, abilityScoreImprovementList)

	if !reflect.DeepEqual(expected, abilityScoreList) {
		t.Errorf("Expected %+v, got %+v", expected, abilityScoreList)
	}
}
