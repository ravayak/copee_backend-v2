package products

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListEquipmentProducts() ([]mysql.EquipmentsProduct, error) {
	ctx := context.Background()
	inventories, err := mysql.QueriesDb.ListEquipmentProducts(ctx)
	return inventories, err
}

func GetEquipmentProduct(productId int32) (mysql.EquipmentsProduct, error) {
	ctx := context.Background()
	inventory, err := mysql.QueriesDb.GetEquipmentProduct(ctx, productId)
	return inventory, err
}

func InsertEquipmentProduct(inventoryParams mysql.CreateEquipmentProductParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateEquipmentProduct(ctx, inventoryParams)
	return res, err
}

func UpdateEquipmentProduct(inventoryParam mysql.UpdateEquipmentProductParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateEquipmentProduct(ctx, inventoryParam)
	return err
}

func DeleteEquipmentProduct(productId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteEquipmentProduct(ctx, productId)
	return err
}
