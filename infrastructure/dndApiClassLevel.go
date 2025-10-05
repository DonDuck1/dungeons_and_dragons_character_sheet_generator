package infrastructure

type DndApiClassLevel struct {
	Level        int                           `json:"level"`
	Spellcasting *DndApiClassLevelSpellcasting `json:"spellcasting"`
	Index        string                        `json:"index"`
	Class        DndApiReference               `json:"class"`
}
