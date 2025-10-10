package infrastructure

type DndApiReference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

func NewDndApiReference(index string, name string, url string) DndApiReference {
	return DndApiReference{Index: index, Name: name, Url: url}
}

func (dndApiReference DndApiReference) GetDeepCopy() DndApiReference {
	return NewDndApiReference(dndApiReference.Index, dndApiReference.Name, dndApiReference.Url)
}
