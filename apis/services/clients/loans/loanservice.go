package loansservice

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/clients/loans"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	clientLoansService clientLoansServiceInterface = &ClientloansService{}
)

type ClientloansService struct{}
type clientLoansServiceInterface interface {
	InsertClientLoan(clientLoanParams mysql.CreateClientLoanParams) (sql.Result, error)
	UpdateClientLoan(clientLoanParams mysql.UpdateClientLoanParams) error
	DeleteClientLoan(clientLoanId int32) error
	GetClientLoan(clientLoanId int32) (mysql.ClientsLoan, error)
	ListClientLoan() ([]mysql.ClientsLoan, error)
}

func (cs *ClientloansService) InsertClientLoan(clientLoanParams mysql.CreateClientLoanParams) (sql.Result, error) {
	return loans.InsertClientLoan(clientLoanParams)
}

func (cs *ClientloansService) UpdateClientLoan(clientLoanParams mysql.UpdateClientLoanParams) error {
	return loans.UpdateClientLoan(clientLoanParams)
}

func (cs *ClientloansService) DeleteClientLoan(clientLoanId int32) error {
	return loans.DeleteClientLoan(clientLoanId)
}

func (cs *ClientloansService) GetClientLoan(clientLoanId int32) (mysql.ClientsLoan, error) {
	return loans.GetClientLoan(clientLoanId)
}

func (cs *ClientloansService) ListClientLoan() ([]mysql.ClientsLoan, error) {
	return loans.ListClientsLoans()
}
