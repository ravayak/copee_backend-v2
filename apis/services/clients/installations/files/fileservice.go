package filesservice

import (
	"database/sql"

	files "github.com/ravayak/copee_backend/apis/domain/clients/installations/files"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	ClientsFilesService clientsFilesServiceInterface = &clientsFilesService{}
)

type clientsFilesService struct{}

type clientsFilesServiceInterface interface {
	InsertClientInstallationFile(clientFileParams mysql.CreateClientInstallationFileParams) (sql.Result, error)
	UpdateClientFile(clientFileParams mysql.UpdateClientInstallationFileParams) error
	DeleteClientFile(clientFileId int32) error
	GetClientFile(clientFileId int32) (mysql.ClientsInstallationsFile, error)
	ListClientFile() ([]mysql.ClientsInstallationsFile, error)
}

func (cs *clientsFilesService) InsertClientInstallationFile(clientFileParams mysql.CreateClientInstallationFileParams) (sql.Result, error) {
	return files.InsertClientInstallationFile(clientFileParams)
}

func (cs *clientsFilesService) UpdateClientFile(clientFileParams mysql.UpdateClientInstallationFileParams) error {
	return files.UpdateClientInstallationFile(clientFileParams)
}

func (cs *clientsFilesService) DeleteClientFile(clientFileId int32) error {
	return files.DeleteClientInstallationFile(clientFileId)
}

func (cs *clientsFilesService) GetClientFile(clientFileId int32) (mysql.ClientsInstallationsFile, error) {
	return files.GetClientInstallationFile(clientFileId)
}

func (cs *clientsFilesService) ListClientFile() ([]mysql.ClientsInstallationsFile, error) {
	return files.ListClientInstallationFile()
}
