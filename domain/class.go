package domain

type Class struct {
	name                    string
	level                   int
	skillProficiencies      []SkillProficiencyName
	classSpellcastingInfo   *ClassSpellcastingInfo
	classWarlockCastingInfo *ClassWarlockCastingInfo
}

func NewClass(name string, level int, skillProficiencies []SkillProficiencyName, classSpellcastingInfo *ClassSpellcastingInfo, classWarlockCastingInfo *ClassWarlockCastingInfo) Class {
	if level < 1 {
		level = 1
	} else if level > 20 {
		level = 20
	}

	return Class{name: name, level: level, skillProficiencies: skillProficiencies, classSpellcastingInfo: classSpellcastingInfo, classWarlockCastingInfo: classWarlockCastingInfo}
}
