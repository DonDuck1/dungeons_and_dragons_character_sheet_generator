package infrastructure

type CsvEquipment struct {
	Name          string
	EquipmentType string
}

func NewCsvEquipment(name string, equipmentType string) CsvEquipment {
	return CsvEquipment{Name: name, EquipmentType: equipmentType}
}
