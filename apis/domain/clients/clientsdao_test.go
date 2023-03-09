package clients

import (
	"database/sql"
	"fmt"
	"testing"

	i "github.com/ravayak/copee_backend/apis/domain/clients/installations"
	"github.com/ravayak/copee_backend/apis/domain/users"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestClientsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Create clients
	resInsertClient, err := InsertClient(dt.CreateClientParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientId, err := resInsertClient.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Client
	client, err := GetClient(int32(clientId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, client.ClientID, int32(clientId))

	clientUpdateParams := mysql.UpdateClientParams{
		ClientLastName:         sql.NullString{String: "Test updated", Valid: true},
		ClientFirstName:        sql.NullString{String: "Test updated", Valid: true},
		ClientEmail:            sql.NullString{String: "testupdated@test.test", Valid: true},
		ClientFiscalYearIncome: sql.NullInt32{Int32: 100000, Valid: true},
		ClientID:               int32(clientId),
	}

	// Update Client
	err = UpdateClient(clientUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated Client
	clients, err := ListClients()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, clients[0].ClientLastName.String, "Test updated")

	// Create installation to test ListClientsByUser !
	// Create user to fullfil Installation Foreign Key constraint
	resCreateUser, err := users.InsertUser(dt.CreateUserParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	userId, err := resCreateUser.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	// Create shared user to fullfil Installation Foreign Key constraint
	resCreateSharedUser, err := users.InsertUser(dt.CreateSharedUserParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	sharedUserId, err := resCreateSharedUser.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Create now client installation
	resInsertInstall, err := i.InsertClientInstallation(dt.GetClientInstallationCreateParams(int32(clientId), int32(userId), int32(sharedUserId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	installId, err := resInsertInstall.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	installation, err := i.GetClientInstallation(int32(installId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, installation.CiClientID.Int32, int32(clientId))

	clientsListByUser, err := ListClientsByUser(int32(userId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, clientsListByUser[0].ClientID, int32(clientId))

	clientIdsListByUser, err := ListClientsIdByUser(int32(userId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientIdsListByUser[0], int32(clientId))

	err = DeleteClient(int32(clientId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = users.DeleteUser(int32(userId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = users.DeleteUser(int32(sharedUserId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
