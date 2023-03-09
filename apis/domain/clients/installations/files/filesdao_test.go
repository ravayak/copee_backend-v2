package files

import (
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/clients/installations"
	itd "github.com/ravayak/copee_backend/apis/domain/clients/installations/tests_dependencies"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestClientInstallationFile(t *testing.T) {
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	clientId, userId, sharedUserId, err := itd.CreateInstallationDepencencies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Create now client installation
	resInsertClientInstallation, err := installations.InsertClientInstallation(dt.GetClientInstallationCreateParams(int32(clientId), int32(userId), int32(sharedUserId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientInstallationId, err := resInsertClientInstallation.LastInsertId()

	// Create Installation Filer
	resInsertClientInstallationFile, err := InsertClientInstallationFile(dt.GetClientInstallationFileParams(int32(clientInstallationId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientInstallationFileId, err := resInsertClientInstallationFile.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Client Installation File
	clientInstallationFile, err := GetClientInstallationFile(int32(clientInstallationFileId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientInstallationFile.CifID, int32(clientInstallationFileId))

	err = UpdateClientInstallationFile(dt.GetUpdateClientInstallationFileParams(int32(clientInstallationFileId), int32(clientInstallationId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Client Installation Files List
	clientInstallationFiles, err := ListClientInstallationFile()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientInstallationFiles[0].CifFileUrl.String, string("Test updated"))

	err = DeleteClientInstallationFile(int32(clientInstallationFileId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Delete Client Installation
	err = installations.DeleteClientInstallation(int32(clientInstallationId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = itd.DeleteInstallationDependencies(int32(clientId), int32(userId), int32(sharedUserId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
