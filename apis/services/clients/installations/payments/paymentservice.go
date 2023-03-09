package paymentsservice

import (
	"database/sql"

	pymt "github.com/ravayak/copee_backend/apis/domain/clients/installations/payments"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	ClientsInstallationPaymentService clientsServiceInstallationPaymentInterface = &clientsInstallationPaymentService{}
)

type clientsInstallationPaymentService struct{}
type clientsServiceInstallationPaymentInterface interface {
	InsertClientInstallationPayment(clientInstallationPaymentParams mysql.CreateClientInstallationPaymentParams) (sql.Result, error)
	UpdateClientInstallationPayment(clientInstallationPaymentParams mysql.UpdateClientInstallationPaymentParams) error
	DeleteClientInstallationPayment(clientInstallationPaymentId int32) error
	GetClientInstallationPayment(clientInstallationPaymentId int32) (mysql.ClientsInstallationsPayment, error)
	ListClientInstallationPayment() ([]mysql.ClientsInstallationsPayment, error)
}

func (cs *clientsInstallationPaymentService) InsertClientInstallationPayment(clientInstallationPaymentParams mysql.CreateClientInstallationPaymentParams) (sql.Result, error) {
	return pymt.InsertClientInstallationPayment(clientInstallationPaymentParams)
}

func (cs *clientsInstallationPaymentService) UpdateClientInstallationPayment(clientInstallationPaymentParams mysql.UpdateClientInstallationPaymentParams) error {
	return pymt.UpdateClientInstallationPayment(clientInstallationPaymentParams)
}

func (cs *clientsInstallationPaymentService) DeleteClientInstallationPayment(clientInstallationPaymentId int32) error {
	return pymt.DeleteClientInstallationPayment(clientInstallationPaymentId)
}

func (cs *clientsInstallationPaymentService) GetClientInstallationPayment(clientInstallationPaymentId int32) (mysql.ClientsInstallationsPayment, error) {
	return pymt.GetClientInstallationPayment(clientInstallationPaymentId)
}

func (cs *clientsInstallationPaymentService) ListClientInstallationPayment() ([]mysql.ClientsInstallationsPayment, error) {
	return pymt.ListClientInstallationPayments()
}
