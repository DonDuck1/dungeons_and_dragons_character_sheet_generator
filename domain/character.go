package domain

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
	proficiencyBonus int,
	abilityScoreList AbilityScoreList,
	skillProficiencyList SkillProficiencyList,
	armorClass int,
	initiative int,
	passivePerception int,
	inventory Inventory,
) *Character {
	character := Character{
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

	return &character
}
