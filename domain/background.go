package domain

type Background struct {
	Name               string
	SkillProficiencies []SkillProficiencyName
}

func NewBackground(name string, skillProficiencies []SkillProficiencyName) Background {
	return Background{Name: name, SkillProficiencies: skillProficiencies}
}
