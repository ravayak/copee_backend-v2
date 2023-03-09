package clients

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListClients() ([]mysql.Client, error) {
	ctx := context.Background()
	clients, err := mysql.QueriesDb.ListClients(ctx)
	return clients, err
}

func ListClientsByUser(userId int32) ([]mysql.Client, error) {
	ctx := context.Background()
	clients, err := mysql.QueriesDb.ListClientsByUser(ctx, sql.NullInt32{Int32: userId, Valid: true})
	return clients, err
}

func ListClientsIdByUser(userId int32) ([]int32, error) {
	ctx := context.Background()
	clientIds, err := mysql.QueriesDb.ListClientsIdByUser(ctx, sql.NullInt32{Int32: userId, Valid: true})
	return clientIds, err
}

func GetClient(clientId int32) (mysql.Client, error) {
	ctx := context.Background()
	client, err := mysql.QueriesDb.GetClient(ctx, clientId)
	return client, err
}

func InsertClient(clientParams mysql.CreateClientParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClient(ctx, clientParams)
	return res, err
}

func UpdateClient(clientParams mysql.UpdateClientParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClient(ctx, clientParams)
	return err
}

func DeleteClient(clientId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClient(ctx, clientId)
	return err
}
