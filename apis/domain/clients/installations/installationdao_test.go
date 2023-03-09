package installations

import (
	"database/sql"
	"fmt"
	"testing"

	itd "github.com/ravayak/copee_backend/apis/domain/clients/installations/tests_dependencies"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestClientsInstallationsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	clientId, userId, sharedUserId, err := itd.CreateInstallationDepencencies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Create now client installation
	resInsertClientInstallation, err := InsertClientInstallation(dt.GetClientInstallationCreateParams(int32(clientId), int32(userId), int32(sharedUserId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientInstallationId, err := resInsertClientInstallation.LastInsertId()

	// Get Client Installation
	clientInstallation, err := GetClientInstallation(int32(clientInstallationId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientInstallation.CiClientID.Int32, int32(clientId))
	assert.Equal(t, clientInstallation.CiUserID.Int32, int32(userId))
	assert.Equal(t, clientInstallation.CiSharedUserID.Int32, int32(sharedUserId))

	clientUpdateParams := mysql.UpdateClientInstallationParams{
		CiClientID:     sql.NullInt32{Int32: int32(clientId), Valid: true},
		CiUserID:       sql.NullInt32{Int32: int32(userId), Valid: true},
		CiSharedUserID: sql.NullInt32{Int32: int32(sharedUserId), Valid: true},
		CiID:           int32(clientInstallationId),
		CiCreationDate: clientInstallation.CiCreationDate,
		CiUpdateDate:   clientInstallation.CiUpdateDate,
		CiStatus:       sql.NullString{String: "Test Updated", Valid: true},
	}

	// Update Client
	err = UpdateClientInstallation(clientUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated Client
	clientsInstallations, err := ListClientsInstallation()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, clientsInstallations[0].CiStatus.String, "Test Updated")

	// Delete Client installation
	err = DeleteClientInstallation(int32(clientInstallationId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = itd.DeleteInstallationDependencies(int32(userId), int32(sharedUserId), int32(clientId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
