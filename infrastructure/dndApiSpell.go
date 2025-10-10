package infrastructure

type DndApiSpell struct {
	Index      string            `json:"index"`
	Name       string            `json:"name"`
	SpellRange string            `json:"range"`
	Level      int               `json:"level"`
	School     DndApiReference   `json:"school"`
	Classes    []DndApiReference `json:"classes"`
}
