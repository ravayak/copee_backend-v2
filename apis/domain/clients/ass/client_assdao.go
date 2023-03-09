package clients

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListClientsAss() ([]mysql.ClientsAss, error) {
	ctx := context.Background()
	clients_ass, err := mysql.QueriesDb.ListClientsAss(ctx)
	return clients_ass, err
}

func GetClientAss(clientId int32) (mysql.ClientsAss, error) {
	ctx := context.Background()
	clients_ass, err := mysql.QueriesDb.GetClientAss(ctx, clientId)
	return clients_ass, err
}

func InsertClientAss(clientParams mysql.CreateClientAssParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientAss(ctx, clientParams)
	return res, err
}

func UpdateClientAss(clientParams mysql.UpdateClientAssParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientAss(ctx, clientParams)
	return err
}

func DeleteClientAss(clientId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientAss(ctx, clientId)
	return err
}
