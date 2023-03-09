package installations

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListClientsInstallation() ([]mysql.ClientsInstallation, error) {
	ctx := context.Background()
	clientsInstallations, err := mysql.QueriesDb.ListClientsInstallation(ctx)
	return clientsInstallations, err
}

func GetClientInstallation(clientInstallationId int32) (mysql.ClientsInstallation, error) {
	ctx := context.Background()
	clientInstallation, err := mysql.QueriesDb.GetClientInstallation(ctx, clientInstallationId)
	return clientInstallation, err
}

func InsertClientInstallation(clientInstallationParams mysql.CreateClientInstallationParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientInstallation(ctx, clientInstallationParams)
	return res, err
}

func UpdateClientInstallation(updateClientInstallationParams mysql.UpdateClientInstallationParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientInstallation(ctx, updateClientInstallationParams)
	return err
}

func DeleteClientInstallation(clientInstallationId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientInstallation(ctx, clientInstallationId)
	return err
}
