package domain

import "math"

type Character struct {
	Name                 string
	Race                 Race
	MainClass            Class
	Background           Background
	ProficiencyBonus     int
	AbilityScoreList     AbilityScoreList
	SkillProficiencyList SkillProficiencyList
	ArmorClass           int
	Initiative           int
	PassivePerception    int
	Inventory            Inventory
}

func NewCharacter(
	name string,
	race Race,
	mainClass Class,
	background Background,
	abilityScoreValueList AbilityScoreValueList,
) Character {
	proficiencyBonus := int(math.Ceil(float64(mainClass.Level)/4)) + 1

	abilityScoreImprovementList := NewAbilityScoreImprovementList(race.AbilityScoreImprovements)
	abilityScoreList := NewAbilityScoreList(abilityScoreValueList, abilityScoreImprovementList)

	skillProficiencies := []SkillProficiencyName{}
	skillProficiencies = append(skillProficiencies, mainClass.SkillProficiencies...)
	skillProficiencies = append(skillProficiencies, background.SkillProficiencies...)
	skillProficiencyList := NewSkillProficiencyList(&abilityScoreList, skillProficiencies, proficiencyBonus)

	inventory := NewEmptyInventory(race.NumberOfHandSlots)

	armorClass := inventory.GetArmorClass(abilityScoreList.Dexterity.Modifier)

	initiative := abilityScoreList.Dexterity.Modifier

	passivePerception := 10 + skillProficiencyList.Perception.Modifier

	return Character{
		Name:                 name,
		Race:                 race,
		MainClass:            mainClass,
		Background:           background,
		ProficiencyBonus:     proficiencyBonus,
		AbilityScoreList:     abilityScoreList,
		SkillProficiencyList: skillProficiencyList,
		ArmorClass:           armorClass,
		Initiative:           initiative,
		PassivePerception:    passivePerception,
		Inventory:            inventory,
	}
}
