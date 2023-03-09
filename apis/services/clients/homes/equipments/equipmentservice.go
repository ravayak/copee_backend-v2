package equipmentsservice

import (
	"database/sql"

	eqpmt "github.com/ravayak/copee_backend/apis/domain/clients/homes/equipments"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	ClientsHomesEquipmentsService clientsHomesEquipmentsServiceInterface = &clientsHomesEquipmentsService{}
)

type clientsHomesEquipmentsService struct{}

type clientsHomesEquipmentsServiceInterface interface {
	InsertClientHomeEquipment(clientHomeEquipmentParams mysql.CreateClientHomeEquipmentParams) (sql.Result, error)
	UpdateClientHomeEquipment(clientHomeEquipmentParams mysql.UpdateClientHomeEquipmentParams) error
	DeleteClientHomeEquipment(clientHomeEquipmentId int32) error
	GetClientHomeEquipmentParams(clientHomeEquipmentId int32) (mysql.ClientsHomesEquipment, error)
	ListClientHomeEquipment() ([]mysql.ClientsHomesEquipment, error)
}

func (cs *clientsHomesEquipmentsService) InsertClientHomeEquipment(clientHomeEquipmentParams mysql.CreateClientHomeEquipmentParams) (sql.Result, error) {
	return eqpmt.InsertClientHomeEquipment(clientHomeEquipmentParams)
}

func (cs *clientsHomesEquipmentsService) UpdateClientHomeEquipment(clientHomeEquipmentParams mysql.UpdateClientHomeEquipmentParams) error {
	return eqpmt.UpdateClientHomeEquipment(clientHomeEquipmentParams)
}

func (cs *clientsHomesEquipmentsService) DeleteClientHomeEquipment(clientHomeEquipmentId int32) error {
	return eqpmt.DeleteClientHomeEquipment(clientHomeEquipmentId)
}

func (cs *clientsHomesEquipmentsService) GetClientHomeEquipmentParams(clientHomeEquipmentId int32) (mysql.ClientsHomesEquipment, error) {
	return eqpmt.GetClientHomeEquipmentParams(clientHomeEquipmentId)
}

func (cs *clientsHomesEquipmentsService) ListClientHomeEquipment() ([]mysql.ClientsHomesEquipment, error) {
	return eqpmt.ListClientHomeEquipments()
}
