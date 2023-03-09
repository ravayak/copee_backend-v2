package catalogue

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListEquipmentCatalogue() ([]mysql.EquipmentsCatalogue, error) {
	ctx := context.Background()
	catalogues, err := mysql.QueriesDb.ListEquipmentCatalogue(ctx)
	return catalogues, err
}

func GetEquipmentCatalogue(catalogueID int32) (mysql.EquipmentsCatalogue, error) {
	ctx := context.Background()
	catalogue, err := mysql.QueriesDb.GetEquipmentCatalogue(ctx, catalogueID)
	return catalogue, err
}

func InsertEquipmentCatalogue(catalogueParams mysql.CreateEquipmentCatalogueParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateEquipmentCatalogue(ctx, catalogueParams)
	return res, err
}

func UpdateEquipmentCatalogue(catalogueParams mysql.UpdateEquipmentCatalogueParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateEquipmentCatalogue(ctx, catalogueParams)
	return err
}

func DeleteEquipmentCatalogue(catalogueId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteEquipmentCatalogue(ctx, catalogueId)
	return err
}
