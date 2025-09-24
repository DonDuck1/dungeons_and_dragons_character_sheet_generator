package domain

type Character struct {
	name                 string
	race                 Race
	mainClass            Class
	background           Background
	proficiencyBonus     int
	abilityScoreList     AbilityScoreList
	skillProficiencyList skillProficiencyList
	armorClass           int
	initiative           int
	passivePerception    int
	inventory            Inventory
}
