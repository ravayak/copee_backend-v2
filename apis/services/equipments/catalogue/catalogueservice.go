package catalogues

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/equipments/catalogue"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	CatalogueService catalogueServiceInterface = &catalogueService{}
)

type catalogueService struct{}
type catalogueServiceInterface interface {
	InsertCatalogue(catalogueParams mysql.CreateEquipmentCatalogueParams) (sql.Result, error)
	UpdateCatalogue(catalogueParams mysql.UpdateEquipmentCatalogueParams) error
	DeleteCatalogue(catalogueId int32) error
	GetCatalogue(catalogueId int32) (mysql.EquipmentsCatalogue, error)
	ListCatalogues() ([]mysql.EquipmentsCatalogue, error)
}

func (is *catalogueService) InsertCatalogue(inventoryParams mysql.CreateEquipmentCatalogueParams) (sql.Result, error) {
	return catalogue.InsertEquipmentCatalogue(inventoryParams)
}

func (is *catalogueService) UpdateCatalogue(inventoryParams mysql.UpdateEquipmentCatalogueParams) error {
	return catalogue.UpdateEquipmentCatalogue(inventoryParams)
}

func (is *catalogueService) DeleteCatalogue(catalogueId int32) error {
	return catalogue.DeleteEquipmentCatalogue(catalogueId)
}

func (is *catalogueService) GetCatalogue(catalogueId int32) (mysql.EquipmentsCatalogue, error) {
	return catalogue.GetEquipmentCatalogue(catalogueId)
}

func (is *catalogueService) ListCatalogues() ([]mysql.EquipmentsCatalogue, error) {
	return catalogue.ListEquipmentCatalogue()
}
