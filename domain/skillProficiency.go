package domain

type SkillProficiency struct {
	Name                      SkillProficiencyName
	Proficient                bool
	TimesProficiencyIsApplied int
	RelatedAbilityScore       *AbilityScore
	Modifier                  int
}

func NewSkillProficiency(name SkillProficiencyName, proficient bool, TimesProficiencyIsApplied int, relatedAbilityScore *AbilityScore, proficiencyBonus int) SkillProficiency {
	modifier := 0

	if proficient {
		modifier += proficiencyBonus
	}

	modifier += relatedAbilityScore.Modifier
	return SkillProficiency{Name: name, Proficient: proficient, RelatedAbilityScore: relatedAbilityScore, Modifier: modifier}
}

func (skillProficiency *SkillProficiency) MakeProficient() {
	skillProficiency.Proficient = true
	skillProficiency.TimesProficiencyIsApplied += 1
}

func (skillProficiency *SkillProficiency) CalculateModifier(proficiencyBonus int) {
	modifier := 0

	if skillProficiency.Proficient {
		modifier += proficiencyBonus
	}

	modifier += skillProficiency.RelatedAbilityScore.Modifier
	skillProficiency.Modifier = modifier
}
