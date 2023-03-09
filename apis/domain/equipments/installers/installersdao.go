package installers

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListEquipmentInstallers() ([]mysql.EquipmentsInstaller, error) {
	ctx := context.Background()
	installers, err := mysql.QueriesDb.ListEquipmentInstallers(ctx)
	return installers, err
}

func GetEquipmentInstaller(installerId int32) (mysql.EquipmentsInstaller, error) {
	ctx := context.Background()
	installer, err := mysql.QueriesDb.GetEquipmentInstaller(ctx, installerId)
	return installer, err
}

func InsertEquipmentInstaller(installerParams mysql.CreateEquipmentInstallerParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateEquipmentInstaller(ctx, installerParams)
	return res, err
}

func UpdateEquipmentInstaller(installerParams mysql.UpdateEquipmentInstallerParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateEquipmentInstaller(ctx, installerParams)
	return err
}

func DeleteEquipmentInstaller(installerId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteEquipmentInstaller(ctx, installerId)
	return err
}
