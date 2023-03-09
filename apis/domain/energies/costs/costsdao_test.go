package energies

import (
	"database/sql"
	"fmt"
	"testing"

	dt "github.com/ravayak/copee_backend/app/data/tests"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestEnergiesCostDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")

	// Create Energies Cost
	resInsertEnergiesCost, err := InsertEnergiesCost(dt.CreateEnergiesParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	energiesCostId, err := resInsertEnergiesCost.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Energies Cost
	energiesCost, err := GetEnergiesCost(int32(energiesCostId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, energiesCost.EncType.String, "Test")

	// Update Energies Cost
	energiesCostUpdateParams := mysql.UpdateEnergiesCostParams{
		EncType: sql.NullString{String: "Test updated", Valid: true},
		EncCost: sql.NullString{String: "123.00", Valid: true},
		EncDate: energiesCost.EncDate,
		EncID:   int32(energiesCostId),
	}

	// Update Energies Cost
	err = UpdateEnergiesCost(energiesCostUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated company
	energiesCosts, err := ListEnergiesCosts()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, energiesCosts[0].EncType.String, "Test updated")

	// Delete Energies Cost
	err = DeleteEnergiesCost(int32(energiesCostId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
