package clientsinstallationsservice

import (
	"database/sql"

	installations "github.com/ravayak/copee_backend/apis/domain/clients/installations"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	ClientsInstallationService clientsServiceInstallationInterface = &clientsInstallationService{}
)

type clientsInstallationService struct{}
type clientsServiceInstallationInterface interface {
	InsertClientInstallation(clientInstallationParams mysql.CreateClientInstallationParams) (sql.Result, error)
	UpdateClientInstallation(clientInstallationParams mysql.UpdateClientInstallationParams) error
	DeleteClientInstallation(clientInstallationId int32) error
	GetClientInstallation(clientInstallationId int32) (mysql.ClientsInstallation, error)
	ListClientInstallation() ([]mysql.ClientsInstallation, error)
}

func (cs *clientsInstallationService) InsertClientInstallation(clientInstallationParams mysql.CreateClientInstallationParams) (sql.Result, error) {
	return installations.InsertClientInstallation(clientInstallationParams)
}

func (cs *clientsInstallationService) UpdateClientInstallation(clientInstallationParams mysql.UpdateClientInstallationParams) error {
	return installations.UpdateClientInstallation(clientInstallationParams)
}

func (cs *clientsInstallationService) DeleteClientInstallation(clientInstallationId int32) error {
	return installations.DeleteClientInstallation(clientInstallationId)
}

func (cs *clientsInstallationService) GetClientInstallation(clientInstallationId int32) (mysql.ClientsInstallation, error) {
	return installations.GetClientInstallation(clientInstallationId)
}

func (cs *clientsInstallationService) ListClientInstallation() ([]mysql.ClientsInstallation, error) {
	return installations.ListClientsInstallation()
}
