package homes

import (
	"database/sql"
	"fmt"
	"testing"

	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestListClientHomes(t *testing.T) {
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	geoId, clientId, err := CreateClientHomesDependencies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resInsertClientHome, err := InsertClientHome(dt.GetClientHomeParams(int32(clientId), int32(geoId)))

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientHomeId, err := resInsertClientHome.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientHome, err := GetClientHomeParams(int32(clientHomeId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientHome.ChID, int32(clientHomeId))

	clientHomeUpdateParams := mysql.UpdateClientHomeParams{
		ChID:               int32(clientHomeId),
		ChClientID:         sql.NullInt32{Int32: int32(clientId), Valid: true},
		ChGeoID:            sql.NullInt32{Int32: int32(geoId), Valid: true},
		ChConstructionYear: sql.NullInt32{Int32: 2000, Valid: true},
		ChArea:             sql.NullInt32{Int32: 100, Valid: true},
		ChResidents:        sql.NullInt32{Int32: 4, Valid: true},
		ChRoofPositionning: sql.NullString{String: "Nord", Valid: true},
		ChRoofSlope:        sql.NullInt32{Int32: 45, Valid: true},
		ChLabel:            sql.NullString{String: "A", Valid: true},
		ChTr:               sql.NullInt32{Int32: 100, Valid: true},
		ChHuc:              sql.NullFloat64{Float64: 100.0, Valid: true},
		ChIsolation:        sql.NullString{String: "Original", Valid: true},
	}

	err = UpdateClientHome(clientHomeUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientHomeList, err := ListClientHome()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, clientHomeList[0].ChIsolation.String, string("Original"))

	clientIds := []int32{int32(clientId)}
	listClientHomesByClientsID, err := ListClientHomesByClients(clientIds)
	assert.Equal(t, listClientHomesByClientsID[0].ChLabel.String, string("A"))

	err = DeleteClientHome(int32(clientHomeId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = DeleteClientHomesDependencies(int32(clientId), int32(geoId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
