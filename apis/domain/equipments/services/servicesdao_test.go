package services

import (
	"database/sql"
	"fmt"
	"testing"

	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestServicesDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	createEqmptServiceParams := mysql.CreateEquipmentServicesParams{
		EsDescription: sql.NullString{String: "Test", Valid: true},
		EsPrice:       sql.NullString{String: "10.00", Valid: true},
	}

	resInsertEquipmentService, err := InsertEquipmentService(createEqmptServiceParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	eqmptServiceId, err := resInsertEquipmentService.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	eqmtService, err := GetEquipmentService(int32(eqmptServiceId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, eqmtService.EsID, int32(eqmptServiceId))

	equipmentServiceUpdatedParams := mysql.UpdateEquipmentServiceParams{
		EsID:          int32(eqmptServiceId),
		EsDescription: sql.NullString{String: "Test updated", Valid: true},
		EsPrice:       sql.NullString{String: "10.00", Valid: true},
	}

	err = UpdateEquipmentService(equipmentServiceUpdatedParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	services, err := ListEquipmentServices()
	assert.Equal(t, services[0].EsDescription.String, "Test updated")

	err = DeleteEquipmentService(int32(eqmptServiceId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
