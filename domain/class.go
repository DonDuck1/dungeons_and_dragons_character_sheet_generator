package domain

type Class struct {
	Name                    ClassName
	Level                   int
	SkillProficiencies      []SkillProficiencyName
	ClassSpellcastingInfo   *ClassSpellcastingInfo
	ClassWarlockCastingInfo *ClassWarlockCastingInfo
}

func NewClass(name ClassName, level int, skillProficiencies []SkillProficiencyName, classSpellcastingInfo *ClassSpellcastingInfo, classWarlockCastingInfo *ClassWarlockCastingInfo) Class {
	if level < 1 {
		level = 1
	} else if level > 20 {
		level = 20
	}

	return Class{Name: name, Level: level, SkillProficiencies: skillProficiencies, ClassSpellcastingInfo: classSpellcastingInfo, ClassWarlockCastingInfo: classWarlockCastingInfo}
}
