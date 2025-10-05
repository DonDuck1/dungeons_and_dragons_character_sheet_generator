package domain

import "fmt"

type AbilityScoreName string

const (
	STRENGTH     AbilityScoreName = "Strength"
	DEXTERITY    AbilityScoreName = "Dexterity"
	CONSTITUTION AbilityScoreName = "Constitution"
	INTELLIGENCE AbilityScoreName = "Intelligence"
	WISDOM       AbilityScoreName = "Wisdom"
	CHARISMA     AbilityScoreName = "Charisma"
)

func AbilityScoreNameFromApiIndex(index string) (AbilityScoreName, error) {
	switch index {
	case "str":
		return STRENGTH, nil
	case "dex":
		return DEXTERITY, nil
	case "con":
		return CONSTITUTION, nil
	case "int":
		return INTELLIGENCE, nil
	case "wis":
		return WISDOM, nil
	case "cha":
		return CHARISMA, nil
	default:
		return CHARISMA, fmt.Errorf("no ability score with index '%s' found", index)
	}
}
