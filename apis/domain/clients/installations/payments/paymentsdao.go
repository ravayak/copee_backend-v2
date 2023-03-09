package payments

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListClientInstallationPayments() ([]mysql.ClientsInstallationsPayment, error) {
	ctx := context.Background()
	clientsInstallationsPayments, err := mysql.QueriesDb.ListClientsInstallationsPayments(ctx)
	return clientsInstallationsPayments, err
}

func GetClientInstallationPayment(clientInstallationPaymentsId int32) (mysql.ClientsInstallationsPayment, error) {
	ctx := context.Background()
	clientsInstallationsPayments, err := mysql.QueriesDb.GetClientInstallationPayment(ctx, clientInstallationPaymentsId)
	return clientsInstallationsPayments, err
}

func InsertClientInstallationPayment(clientInstallationsPaymentParams mysql.CreateClientInstallationPaymentParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientInstallationPayment(ctx, clientInstallationsPaymentParams)
	return res, err
}

func UpdateClientInstallationPayment(updateClientInstallationFileParams mysql.UpdateClientInstallationPaymentParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientInstallationPayment(ctx, updateClientInstallationFileParams)
	return err
}

func DeleteClientInstallationPayment(clientInstallationPaymentsId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientInstallationPayment(ctx, clientInstallationPaymentsId)
	return err
}
