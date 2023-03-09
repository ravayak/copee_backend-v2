package equipment

import (
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/clients/homes"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestListClientEquipment(t *testing.T) {
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	geoId, clientId, err := homes.CreateClientHomesDependencies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resInsertClientHome, err := homes.InsertClientHome(dt.GetClientHomeParams(int32(clientId), int32(geoId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	homeId, err := resInsertClientHome.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resInsertClientHomeEquipment, err := InsertClientHomeEquipment(dt.GetClientHomeEquipmentParams(int32(homeId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	homeEqmtId, err := resInsertClientHomeEquipment.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientHomeEquipment, err := GetClientHomeEquipmentParams(int32(homeEqmtId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientHomeEquipment.CheID, int32(homeEqmtId))

	clientUpdateParams := dt.GetClientHomeEquipmentUpdateParams(int32(homeEqmtId), int32(homeId))
	err = UpdateClientHomeEquipment(clientUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientHomesEquipmentList, err := ListClientHomeEquipments()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, clientHomesEquipmentList[0].CheEquipmentType.String, "Test updated")

	clientHomeEquipmentList, err := ListClientHomeEquipmentsByHomes([]int32{int32(homeId)})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, clientHomeEquipmentList[0].CheID, int32(homeEqmtId))

	err = DeleteClientHomeEquipment(int32(homeEqmtId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = homes.DeleteClientHomesDependencies(int32(clientId), int32(geoId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
