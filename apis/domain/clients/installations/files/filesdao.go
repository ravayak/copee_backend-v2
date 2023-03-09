package files

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListClientInstallationFile() ([]mysql.ClientsInstallationsFile, error) {
	ctx := context.Background()
	clientsInstallationsFiles, err := mysql.QueriesDb.ListClientsInstallationFiles(ctx)
	return clientsInstallationsFiles, err
}

func GetClientInstallationFile(clientInstallationFilesId int32) (mysql.ClientsInstallationsFile, error) {
	ctx := context.Background()
	clientsInstallationsFiles, err := mysql.QueriesDb.GetClientInstallationFile(ctx, clientInstallationFilesId)
	return clientsInstallationsFiles, err
}

func InsertClientInstallationFile(clientInstallationFilesParams mysql.CreateClientInstallationFileParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientInstallationFile(ctx, clientInstallationFilesParams)
	return res, err
}

func UpdateClientInstallationFile(updateClientInstallationFileParams mysql.UpdateClientInstallationFileParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientInstallationFile(ctx, updateClientInstallationFileParams)
	return err
}

func DeleteClientInstallationFile(clientInstallationFilesId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientInstallationFile(ctx, clientInstallationFilesId)
	return err
}
