package designationsservice

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/equipments/installers"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	InstallersService installersServiceInterface = &installersService{}
)

type installersService struct{}
type installersServiceInterface interface {
	InsertInstallers(designParams mysql.CreateEquipmentInstallerParams) (sql.Result, error)
	UpdateInstallers(designParams mysql.UpdateEquipmentInstallerParams) error
	DeleteInstallers(designId int32) error
	GetInstallers(designId int32) (mysql.EquipmentsInstaller, error)
	ListInstallers() ([]mysql.EquipmentsInstaller, error)
}

func (is *installersService) InsertInstallers(installerParams mysql.CreateEquipmentInstallerParams) (sql.Result, error) {
	return installers.InsertEquipmentInstaller(installerParams)
}

func (is *installersService) UpdateInstallers(installerParams mysql.UpdateEquipmentInstallerParams) error {
	return installers.UpdateEquipmentInstaller(installerParams)
}

func (is *installersService) DeleteInstallers(installerId int32) error {
	return installers.DeleteEquipmentInstaller(installerId)
}

func (is *installersService) GetInstallers(installerId int32) (mysql.EquipmentsInstaller, error) {
	return installers.GetEquipmentInstaller(installerId)
}

func (is *installersService) ListInstallers() ([]mysql.EquipmentsInstaller, error) {
	return installers.ListEquipmentInstallers()
}
