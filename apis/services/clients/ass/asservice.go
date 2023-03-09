package clientasssservice

import (
	"database/sql"

	ass "github.com/ravayak/copee_backend/apis/domain/clients/ass"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	ClientAssService clientAssServiceInterface = &clientAssService{}
)

type clientAssService struct{}
type clientAssServiceInterface interface {
	InsertClientAss(clientAssParams mysql.CreateClientAssParams) (sql.Result, error)
	UpdateClientAss(clientAssParams mysql.UpdateClientAssParams) error
	DeleteClientAss(clientAssId int32) error
	GetClientAss(clientAssId int32) (mysql.ClientsAss, error)
	ListClientAss() ([]mysql.ClientsAss, error)
}

func (cs *clientAssService) InsertClientAss(clientAssParams mysql.CreateClientAssParams) (sql.Result, error) {
	return ass.InsertClientAss(clientAssParams)
}

func (cs *clientAssService) UpdateClientAss(clientAssParams mysql.UpdateClientAssParams) error {
	return ass.UpdateClientAss(clientAssParams)
}

func (cs *clientAssService) DeleteClientAss(clientAssId int32) error {
	return ass.DeleteClientAss(clientAssId)
}

func (cs *clientAssService) GetClientAss(clientAssId int32) (mysql.ClientsAss, error) {
	return ass.GetClientAss(clientAssId)
}

func (cs *clientAssService) ListClientAss() ([]mysql.ClientsAss, error) {
	return ass.ListClientsAss()
}
