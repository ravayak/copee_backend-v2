package loans

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListClientsLoans() ([]mysql.ClientsLoan, error) {
	ctx := context.Background()
	clientsLoans, err := mysql.QueriesDb.ListClientsLoans(ctx)
	return clientsLoans, err
}

func GetClientLoan(clientLoanId int32) (mysql.ClientsLoan, error) {
	ctx := context.Background()
	clientLoan, err := mysql.QueriesDb.GetClientLoan(ctx, clientLoanId)
	return clientLoan, err
}

func InsertClientLoan(clientLoanParams mysql.CreateClientLoanParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientLoan(ctx, clientLoanParams)
	return res, err
}

func UpdateClientLoan(updateClientLoanParams mysql.UpdateClientLoanParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientLoan(ctx, updateClientLoanParams)
	return err
}

func DeleteClientLoan(clientLoanId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientLoan(ctx, clientLoanId)
	return err
}
