package domain

import "fmt"

type Background struct {
	name               BackgroundName
	skillProficiencies []SkillProficiencyName
}

func NewBackground(name BackgroundName) (Background, error) {
	switch name {
	case ACOLYTE:
		skillProficiencies := []SkillProficiencyName{INSIGHT, RELIGION}
		return Background{name: name, skillProficiencies: skillProficiencies}, nil
	default:
		skillProficiencies := []SkillProficiencyName{}
		err := fmt.Errorf("Unknown background provided: %s", name)
		return Background{name: name, skillProficiencies: skillProficiencies}, err
	}
}
