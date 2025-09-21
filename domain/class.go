package domain

type Class struct {
	name                string
	skillProficiencies  []SkillProficiencyName
	spellcastingClass   bool
	spellsKnownPerLevel [20]int
	spellSlotsPerLevel  [20][9]int
}
