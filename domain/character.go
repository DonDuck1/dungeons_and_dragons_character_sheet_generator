package domain

type Character struct {
	name                 string
	race                 Race
	mainClass            Class
	level                int
	background           Background
	proficiencyBonus     int
	abilityScoreList     AbilityScoreList
	skillProficiencyList skillProficiencyList
	armorClass           int
	initiative           int
	passivePerception    int
	inventory            Inventory
	spellList            SpellList
	spellcastingAbility  AbilityScoreName
	spellSaveDC          int
	spellAttackBonus     int
}
