package clientsservice

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ravayak/copee_backend/apis/domain/clients"
	is "github.com/ravayak/copee_backend/apis/services/clients/installations"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	ClientsService clientsServiceInterface = &clientsService{}
)

type clientsService struct{}
type clientsServiceInterface interface {
	InsertClient(clientParams mysql.CreateClientParams, userId int32) (*clients.Client, error)
	UpdateClient(clientParams mysql.UpdateClientParams) error
	DeleteClient(clientId int32) error
	GetClient(clientId int32) (*clients.Client, error)
	ListClients() ([]*clients.Client, error)
	ListClientsByUser(userId int32) ([]*clients.Client, error)
	ListClientsIdByUser(userId int32) ([]int32, error)
}

func (cs *clientsService) InsertClient(clientParams mysql.CreateClientParams, userId int32) (*clients.Client, error) {
	// userLogin, err := ctxt.GetUserLoginFromContext(c)
	// if err != nil {
	// 	return nil, err
	// }

	res, err := clients.InsertClient(clientParams)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	client, err := cs.GetClient(int32(id))
	if err != nil {
		return nil, err
	}

	is.ClientsInstallationService.InsertClientInstallation(mysql.CreateClientInstallationParams{
		CiClientID:     sql.NullInt32{Int32: client.ClientID, Valid: true},
		CiUserID:       sql.NullInt32{Int32: userId, Valid: true},
		CiCreationDate: sql.NullTime{Time: time.Now(), Valid: true},
		CiUpdateDate:   sql.NullTime{Time: time.Now(), Valid: true},
		CiComments:     sql.NullString{String: fmt.Sprintf("Installation vierge créée suite à la création du client n°%d", client.ClientID), Valid: true},
		CiStatus:       sql.NullString{String: "En attente", Valid: true},
	})

	return client, nil
}

func (cs *clientsService) UpdateClient(clientParams mysql.UpdateClientParams) error {
	return clients.UpdateClient(clientParams)
}

func (cs *clientsService) DeleteClient(clientId int32) error {
	return clients.DeleteClient(clientId)
}

func (cs *clientsService) GetClient(clientId int32) (*clients.Client, error) {
	client, err := clients.GetClient(clientId)
	if err != nil {
		return nil, err
	}
	return convertMysqlClientToClient(client), nil
}

func (cs *clientsService) ListClients() ([]*clients.Client, error) {
	mysqlClients, err := clients.ListClients()
	if err != nil {
		return nil, err
	}
	clientsRes := convertMysqlClientListToClientList(mysqlClients)
	return clientsRes, nil
}

func (cs *clientsService) ListClientsByUser(userId int32) ([]*clients.Client, error) {
	mysqlClients, err := clients.ListClientsByUser(userId)
	if err != nil {
		return nil, err
	}
	clientsRes := convertMysqlClientListToClientList(mysqlClients)
	return clientsRes, nil
}

func (cs *clientsService) ListClientsIdByUser(userId int32) ([]int32, error) {
	return clients.ListClientsIdByUser(userId)
}

// Utility functions to convert from mysql to clients.Client in response
func convertMysqlClientToClient(client mysql.Client) *clients.Client {
	return &clients.Client{
		ClientID:               client.ClientID,
		ClientLastName:         client.ClientLastName.String,
		ClientFirstName:        client.ClientFirstName.String,
		ClientEmail:            client.ClientEmail.String,
		ClientPhone:            client.ClientPhone.String,
		ClientFiscalYearIncome: client.ClientFiscalYearIncome.Int32,
		ClientDateCreated:      client.ClientDateCreated.Time,
	}
}

func convertMysqlClientListToClientList(mysqlClients []mysql.Client) []*clients.Client {
	clientsResponse := []*clients.Client{}
	for _, client := range mysqlClients {
		clientsResponse = append(clientsResponse, convertMysqlClientToClient(client))
	}
	return clientsResponse
}
