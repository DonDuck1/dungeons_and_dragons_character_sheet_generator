package domain

type Background struct {
	name               string
	skillProficiencies []SkillProficiencyName
}

func NewBackground(name string, skillProficiencies []SkillProficiencyName) Background {
	return Background{name: name, skillProficiencies: skillProficiencies}
}
