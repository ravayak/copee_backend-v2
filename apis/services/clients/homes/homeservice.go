package clientsservice

import (
	"database/sql"

	homes "github.com/ravayak/copee_backend/apis/domain/clients/homes"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	ClientsHomesService clientsHomesServiceInterface = &clientsHomesService{}
)

type clientsHomesService struct{}
type clientsHomesServiceInterface interface {
	InsertClientHome(clientHomeParams mysql.CreateClientHomeParams) (sql.Result, error)
	UpdateClientHome(clientHomeParams mysql.UpdateClientHomeParams) error
	DeleteClientHome(clientHomeId int32) error
	GetClientHomeParams(clientHomeId int32) (mysql.ClientsHome, error)
	ListClientHome() ([]mysql.ClientsHome, error)
}

func (cs *clientsHomesService) InsertClientHome(clientHomeParams mysql.CreateClientHomeParams) (sql.Result, error) {
	return homes.InsertClientHome(clientHomeParams)
}

func (cs *clientsHomesService) UpdateClientHome(clientHomeParams mysql.UpdateClientHomeParams) error {
	return homes.UpdateClientHome(clientHomeParams)
}

func (cs *clientsHomesService) DeleteClientHome(clientHomeId int32) error {
	return homes.DeleteClientHome(clientHomeId)
}

func (cs *clientsHomesService) GetClientHomeParams(clientHomeId int32) (mysql.ClientsHome, error) {
	return homes.GetClientHomeParams(clientHomeId)
}

func (cs *clientsHomesService) ListClientHome() ([]mysql.ClientsHome, error) {
	return homes.ListClientHome()
}
