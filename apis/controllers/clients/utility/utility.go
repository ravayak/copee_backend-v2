package clientscontroller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	clients "github.com/ravayak/copee_backend/apis/domain/clients"
	data "github.com/ravayak/copee_backend/apis/domain/clients/data"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

// setClientParams sets the client params for create and update
func SetClientParams(clientParamsType string, clientToEdit clients.Client, c *gin.Context) (interface{}, error) {
	if err := c.ShouldBindJSON(&clientToEdit); err != nil {
		return nil, err
	}
	var clientParams interface{}
	if clientParamsType == data.CREATE_CLIENT_PARAMS {
		clientParams = mysql.CreateClientParams{
			ClientLastName:         sql.NullString{String: clientToEdit.ClientLastName, Valid: true},
			ClientFirstName:        sql.NullString{String: clientToEdit.ClientFirstName, Valid: true},
			ClientEmail:            sql.NullString{String: clientToEdit.ClientEmail, Valid: true},
			ClientPhone:            sql.NullString{String: clientToEdit.ClientPhone, Valid: true},
			ClientFiscalYearIncome: sql.NullInt32{Int32: clientToEdit.ClientFiscalYearIncome, Valid: true},
		}
	}

	if clientParamsType == data.UPDATE_CLIENT_PARAMS {
		clientParams = mysql.UpdateClientParams{
			ClientLastName:         sql.NullString{String: clientToEdit.ClientLastName, Valid: true},
			ClientFirstName:        sql.NullString{String: clientToEdit.ClientFirstName, Valid: true},
			ClientEmail:            sql.NullString{String: clientToEdit.ClientEmail, Valid: true},
			ClientPhone:            sql.NullString{String: clientToEdit.ClientPhone, Valid: true},
			ClientFiscalYearIncome: sql.NullInt32{Int32: clientToEdit.ClientFiscalYearIncome, Valid: true},
			ClientID:               clientToEdit.ClientID,
		}
	}

	return clientParams, nil
}
