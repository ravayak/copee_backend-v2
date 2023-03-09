package designations

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListEquipmentDesignations() ([]mysql.EquipmentsDesignation, error) {
	ctx := context.Background()
	designations, err := mysql.QueriesDb.ListEquipmentDesignations(ctx)
	return designations, err
}

func GetEquipmentDesignation(designationId int32) (mysql.EquipmentsDesignation, error) {
	ctx := context.Background()
	designation, err := mysql.QueriesDb.GetEquipmentDesignation(ctx, designationId)
	return designation, err
}

func InsertEquipmentDesignation(designationParams mysql.CreateEquipmentDesignationsParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateEquipmentDesignations(ctx, designationParams)
	return res, err
}

func UpdateEquipmentDesignation(designationParams mysql.UpdateEquipmentDesignationParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateEquipmentDesignation(ctx, designationParams)
	return err
}

func DeleteEquipmentDesignation(categoryId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteEquipmentDesignation(ctx, categoryId)
	return err
}
