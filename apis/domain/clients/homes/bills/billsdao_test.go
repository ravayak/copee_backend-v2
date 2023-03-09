package bills

import (
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/clients/homes"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestListClientHomesBills(t *testing.T) {
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Create homes
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

	// Create home bills
	resInsertClientHomeBills, err := InsertClientHomesBill(dt.GetClientHomeParamsBillsParams(int32(homeId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientHomeBillsId, err := resInsertClientHomeBills.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientHomesBills, err := GetClientHomeParamsBill(int32(clientHomeBillsId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientHomesBills.ChbID, int32(clientHomeBillsId))

	clientHomeBillsUpdateParams := dt.GetUpdateClientHomeBillsParams(int32(homeId), int32(clientHomeBillsId))

	err = UpdateClientHomesBill(clientHomeBillsUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientHomesBillsList, err := ListClientHomesBills()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, clientHomesBillsList[0].ChbElectricity.Int32, int32(60))

	clientHomeBillList, err := ListClientHomeBillsByHomes([]int32{int32(homeId)})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientHomeBillList[0].ChbID, int32(clientHomeBillsId))

	err = DeleteClientHomesBill(int32(clientHomeBillsId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = homes.DeleteClientHome(int32(homeId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
