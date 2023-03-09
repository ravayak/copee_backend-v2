package designationsservice

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/equipments/designations"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	DesignationsService designationsServiceInterface = &designationsService{}
)

type designationsService struct{}
type designationsServiceInterface interface {
	InsertDesigation(designParams mysql.CreateEquipmentDesignationsParams) (sql.Result, error)
	UpdateDesigation(designParams mysql.UpdateEquipmentDesignationParams) error
	DeleteDesigation(designId int32) error
	GetDesigation(designId int32) (mysql.EquipmentsDesignation, error)
	ListCategories() ([]mysql.EquipmentsDesignation, error)
}

func (is *designationsService) InsertDesigation(designParams mysql.CreateEquipmentDesignationsParams) (sql.Result, error) {
	return designations.InsertEquipmentDesignation(designParams)
}

func (is *designationsService) UpdateDesigation(designParams mysql.UpdateEquipmentDesignationParams) error {
	return designations.UpdateEquipmentDesignation(designParams)
}

func (is *designationsService) DeleteDesigation(designId int32) error {
	return designations.DeleteEquipmentDesignation(designId)
}

func (is *designationsService) GetDesigation(designId int32) (mysql.EquipmentsDesignation, error) {
	return designations.GetEquipmentDesignation(designId)
}

func (is *designationsService) ListCategories() ([]mysql.EquipmentsDesignation, error) {
	return designations.ListEquipmentDesignations()
}
