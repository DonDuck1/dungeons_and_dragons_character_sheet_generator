package domain

type Background struct {
	skillProficiencies []SkillProficiencyName
}

func NewBackground(skillProficiencies []SkillProficiencyName) Background {
	return Background{skillProficiencies: skillProficiencies}
}
