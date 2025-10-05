package infrastructure

type DndApiReferenceList struct {
	Count   int               `json:"count"`
	Results []DndApiReference `json:"results"`
}
