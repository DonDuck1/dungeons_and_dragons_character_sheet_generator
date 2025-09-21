package domain

type Spell struct {
	name       string
	level      int
	classes    []ClassName
	school     string
	spellRange string
}
