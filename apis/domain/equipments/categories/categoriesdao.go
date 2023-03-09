package categories

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListEquipmentCategories() ([]mysql.EquipmentsCategory, error) {
	ctx := context.Background()
	categories, err := mysql.QueriesDb.ListCategories(ctx)
	return categories, err
}

func GetEquipmentCategory(catalogueID int32) (mysql.EquipmentsCategory, error) {
	ctx := context.Background()
	category, err := mysql.QueriesDb.GetCategory(ctx, catalogueID)
	return category, err
}

func InsertEquipmentCategory(categoryParams mysql.CreateCategoryParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateCategory(ctx, categoryParams)
	return res, err
}

func UpdateEquipmentCategory(categoryParams mysql.UpdateCategoryParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateCategory(ctx, categoryParams)
	return err
}

func DeleteEquipmentCategory(categoryId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteCategory(ctx, categoryId)
	return err
}
