package services

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListEquipmentServices() ([]mysql.EquipmentsService, error) {
	ctx := context.Background()
	services, err := mysql.QueriesDb.ListEquipmentServices(ctx)
	return services, err
}

func GetEquipmentService(productId int32) (mysql.EquipmentsService, error) {
	ctx := context.Background()
	service, err := mysql.QueriesDb.GetEquipmentService(ctx, productId)
	return service, err
}

func InsertEquipmentService(serviceParams mysql.CreateEquipmentServicesParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateEquipmentServices(ctx, serviceParams)
	return res, err
}

func UpdateEquipmentService(serviceParams mysql.UpdateEquipmentServiceParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateEquipmentService(ctx, serviceParams)
	return err
}

func DeleteEquipmentService(productId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteEquipmentService(ctx, productId)
	return err
}
