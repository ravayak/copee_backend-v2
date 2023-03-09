package billsservice

import (
	"database/sql"

	bills "github.com/ravayak/copee_backend/apis/domain/clients/homes/bills"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	ClientsHomesBillsService clientsHomesBillsServiceInterface = &clientsHomesBillsService{}
)

type clientsHomesBillsService struct{}

type clientsHomesBillsServiceInterface interface {
	InsertClientHomeBill(clientHomeBillParams mysql.CreateClientHomeBillParams) (sql.Result, error)
	UpdateClientHomeBill(clientHomeBillParams mysql.UpdateClientHomeBillParams) error
	DeleteClientHomeBill(clientHomeBillId int32) error
	GetClientHomeParamsBill(clientHomeBillId int32) (mysql.ClientsHomesBill, error)
	ListClientHomeBill() ([]mysql.ClientsHomesBill, error)
}

func (cs *clientsHomesBillsService) InsertClientHomeBill(clientHomeBillParams mysql.CreateClientHomeBillParams) (sql.Result, error) {
	return bills.InsertClientHomesBill(clientHomeBillParams)
}

func (cs *clientsHomesBillsService) UpdateClientHomeBill(clientHomeBillParams mysql.UpdateClientHomeBillParams) error {
	return bills.UpdateClientHomesBill(clientHomeBillParams)
}

func (cs *clientsHomesBillsService) DeleteClientHomeBill(clientHomeBillId int32) error {
	return bills.DeleteClientHomesBill(clientHomeBillId)
}

func (cs *clientsHomesBillsService) GetClientHomeParamsBill(clientHomeBillId int32) (mysql.ClientsHomesBill, error) {
	return bills.GetClientHomeParamsBill(clientHomeBillId)
}

func (cs *clientsHomesBillsService) ListClientHomeBill() ([]mysql.ClientsHomesBill, error) {
	return bills.ListClientHomesBills()
}
