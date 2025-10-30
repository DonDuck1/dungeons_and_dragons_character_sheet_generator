package infrastructure

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type CsvEquipmentRepository struct {
	filepath      string
	equipmentList *[]CsvEquipment
}

const (
	NO_EQUIPMENT_WITH_TYPE string = "could not find equipment of type '%s'"
)

func NewCsvEquipmentRepository(filepath string) (*CsvEquipmentRepository, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	equipmentList := []CsvEquipment{}

	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if len(record) < 2 {
			continue
		}

		equipment := NewCsvEquipment(strings.TrimSpace(record[0]), strings.TrimSpace(record[1]))
		equipmentList = append(equipmentList, equipment)
	}

	return &CsvEquipmentRepository{
		filepath:      filepath,
		equipmentList: &equipmentList,
	}, nil
}

func (csvEquipmentRepository CsvEquipmentRepository) GetAll() *[]CsvEquipment {
	return csvEquipmentRepository.equipmentList
}

func (csvEquipmentRepository CsvEquipmentRepository) GetByEquipmentType(equipmentType string) (*[]CsvEquipment, error) {
	relevantEquipmentList := []CsvEquipment{}

	for _, equipment := range *csvEquipmentRepository.equipmentList {
		if strings.EqualFold(equipment.EquipmentType, equipmentType) {
			relevantEquipmentList = append(relevantEquipmentList, equipment)
		}
	}

	if len(relevantEquipmentList) != 0 {
		return &relevantEquipmentList, nil
	}

	err := fmt.Errorf(NO_EQUIPMENT_WITH_TYPE, equipmentType)
	return nil, err
}
