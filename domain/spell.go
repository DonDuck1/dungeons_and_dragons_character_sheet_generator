package domain

type Spell struct {
	Name       string
	Level      int
	Classes    []string
	School     string
	SpellRange int
	Prepared   bool
}

func NewSpell(name string, level int, classes []string, school string, spellRange int, prepared bool) Spell {
	return Spell{Name: name, Level: level, Classes: classes, School: school, SpellRange: spellRange, Prepared: prepared}
}
