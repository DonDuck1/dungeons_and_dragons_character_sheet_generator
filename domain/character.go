package domain

type Character struct {
	Name                 string
	Race                 Race
	Class                Class
	Background           Background
	ProficiencyBonus     int
	AbilityScoreList     AbilityScoreList
	SkillProficiencyList SkillProficiencyList
	ArmorClass           int
	Initiative           int
	PassivePerception    int
	Inventory            Inventory
	MaxHitPoints         int
}

func NewCharacter(
	name string,
	race Race,
	class Class,
	background Background,
	proficiencyBonus int,
	abilityScoreList AbilityScoreList,
	skillProficiencyList SkillProficiencyList,
	armorClass int,
	initiative int,
	passivePerception int,
	inventory Inventory,
	maxHitPoints int,
) *Character {
	character := Character{
		Name:                 name,
		Race:                 race,
		Class:                class,
		Background:           background,
		ProficiencyBonus:     proficiencyBonus,
		AbilityScoreList:     abilityScoreList,
		SkillProficiencyList: skillProficiencyList,
		ArmorClass:           armorClass,
		Initiative:           initiative,
		PassivePerception:    passivePerception,
		Inventory:            inventory,
		MaxHitPoints:         maxHitPoints,
	}

	return &character
}
