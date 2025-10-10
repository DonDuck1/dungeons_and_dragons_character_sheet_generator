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

func (spell Spell) CanBeUsedByClass(className ClassName) bool {
	for _, validClassName := range spell.Classes {
		if className == validClassName {
			return true
		}
	}

	return false
}
