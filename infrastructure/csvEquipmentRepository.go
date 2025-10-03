package infrastructure

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type CsvEquipmentRepository struct {
	filepath      string
	equipmentList *[]domain.CsvEquipment
}

func NewCsvEquipmentRepository(filepath string) (*CsvEquipmentRepository, error) {
	_, err := os.Stat(filepath)
	if !(err == nil) {
		return nil, err
	}

	file, err := os.Open(filepath)
	if !(err == nil) {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	equipmentList := []domain.CsvEquipment{}

	_, err = reader.Read()
	if !(err == nil) {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if !(err == nil) {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if len(record) < 2 {
			continue
		}

		equipment := domain.NewCsvEquipment(strings.TrimSpace(record[0]), strings.TrimSpace(record[1]))
		equipmentList = append(equipmentList, equipment)
	}

	return &CsvEquipmentRepository{
		filepath:      filepath,
		equipmentList: &equipmentList,
	}, nil
}

func (csvEquipmentRepository CsvEquipmentRepository) GetAll() *[]domain.CsvEquipment {
	return csvEquipmentRepository.equipmentList
}

func (csvEquipmentRepository CsvEquipmentRepository) GetByEquipmentType(equipmentType string) (*[]domain.CsvEquipment, error) {
	relevantEquipmentList := []domain.CsvEquipment{}

	for _, equipment := range *csvEquipmentRepository.equipmentList {
		if strings.EqualFold(equipment.EquipmentType, equipmentType) {
			relevantEquipmentList = append(relevantEquipmentList, equipment)
		}
	}

	if !(len(relevantEquipmentList) == 0) {
		return &relevantEquipmentList, nil
	}

	err := fmt.Errorf("could not find equipment of type '%s'", equipmentType)
	return nil, err
}
