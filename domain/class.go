package domain

import "fmt"

type Class struct {
	name                    ClassName
	level                   int
	skillProficiencies      []SkillProficiencyName
	classSpellcastingInfo   *ClassSpellcastingInfo
	classWarlockCastingInfo *ClassWarlockCastingInfo
}

func NewClass(name ClassName, level int, abilityScoreList AbilityScoreList, proficiencyBonus int) (Class, error) {
	if !(level >= 1 && level <= 20) {
		skillProficiencies := []SkillProficiencyName{}
		err := fmt.Errorf("Unvalid level provided: %d", level)

		return Class{name: name, level: level, skillProficiencies: skillProficiencies, classSpellcastingInfo: nil, classWarlockCastingInfo: nil}, err
	}

	switch name {
	case BARBARIAN:
		skillProficiencies := []SkillProficiencyName{ANIMAL_HANDLING, ATHLETICS}

	case BARD:

	case CLERIC:

	case DRUID:

	case FIGHTER:

	case MONK:

	case PALADIN:

	case RANGER:

	case ROGUE:

	case SORCERER:

	case WARLOCK:

	case WIZARD:

	default:
		skillProficiencies := []SkillProficiencyName{}
		err := fmt.Errorf("Unknown class provided: %s", name)

		return Class{name: name, level: level, skillProficiencies: skillProficiencies, classSpellcastingInfo: nil, classWarlockCastingInfo: nil}, err
	}
}
