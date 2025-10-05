package domain

type Spell struct {
	Name       string
	Level      int
	Classes    []ClassName
	School     string
	SpellRange string
	Prepared   bool
}

func NewSpell(name string, level int, classes []ClassName, school string, spellRange string, prepared bool) Spell {
	return Spell{Name: name, Level: level, Classes: classes, School: school, SpellRange: spellRange, Prepared: prepared}
}
