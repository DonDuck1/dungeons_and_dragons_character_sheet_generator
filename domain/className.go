package domain

import "fmt"

type ClassName string

const (
	BARBARIAN ClassName = "Barbarian"
	BARD      ClassName = "Bard"
	CLERIC    ClassName = "Cleric"
	DRUID     ClassName = "Druid"
	FIGHTER   ClassName = "Fighter"
	MONK      ClassName = "Monk"
	PALADIN   ClassName = "Paladin"
	RANGER    ClassName = "Ranger"
	ROGUE     ClassName = "Rogue"
	SORCERER  ClassName = "Sorcerer"
	WARLOCK   ClassName = "Warlock"
	WIZARD    ClassName = "Wizard"
)

func ClassNameFromApiIndex(index string) (ClassName, error) {
	switch index {
	case "barbarian":
		return BARBARIAN, nil
	case "bard":
		return BARD, nil
	case "cleric":
		return CLERIC, nil
	case "druid":
		return DRUID, nil
	case "fighter":
		return FIGHTER, nil
	case "monk":
		return MONK, nil
	case "paladin":
		return PALADIN, nil
	case "ranger":
		return RANGER, nil
	case "rogue":
		return ROGUE, nil
	case "sorcerer":
		return SORCERER, nil
	case "warlock":
		return WARLOCK, nil
	case "wizard":
		return WIZARD, nil
	default:
		return WIZARD, fmt.Errorf("no class with index '%s' found", index)
	}
}
