package infrastructure

type DndApiClassLevel struct {
	Level        int                           `json:"level"`
	Spellcasting *DndApiClassLevelSpellcasting `json:"spellcasting"`
	Index        string                        `json:"index"`
	Class        DndApiReference               `json:"class"`
}

func NewDndApiClassLevel(level int, spellcasting *DndApiClassLevelSpellcasting, index string, class DndApiReference) DndApiClassLevel {
	return DndApiClassLevel{
		Level:        level,
		Spellcasting: spellcasting,
		Index:        index,
		Class:        class,
	}
}

func (dndApiClassLevel DndApiClassLevel) GetDeepCopy() DndApiClassLevel {
	var deepCopiedSpellcasting *DndApiClassLevelSpellcasting
	if dndApiClassLevel.Spellcasting != nil {
		value := dndApiClassLevel.Spellcasting.GetDeepCopy()
		deepCopiedSpellcasting = &value
	}

	return NewDndApiClassLevel(dndApiClassLevel.Level, deepCopiedSpellcasting, dndApiClassLevel.Index, dndApiClassLevel.Class)
}
