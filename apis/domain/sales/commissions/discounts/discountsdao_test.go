package discounts

import (
	"database/sql"
	"fmt"
	"testing"

	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCommissionsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	createSalesParams := mysql.CreateSalesCommissionDiscountParams{
		ScDiscountPct: sql.NullFloat64{Float64: 0.1, Valid: true},
		ScComPct:      sql.NullFloat64{Float64: 0.1, Valid: true},
	}

	resCd, err := InsertSalesCommissionsDiscount(createSalesParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	cdId, err := resCd.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	cd, err := GetSalesCommissionsDiscount(int32(cdId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, cd.ScID, int32(cdId))

	updateSalesParams := mysql.UpdateSalesCommissionDiscountParams{
		ScID:          int32(cdId),
		ScDiscountPct: sql.NullFloat64{Float64: 0.2, Valid: true},
		ScComPct:      sql.NullFloat64{Float64: 0.2, Valid: true},
	}

	err = UpdateSalesCommissionsDiscount(updateSalesParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	cdList, err := ListSalesCommissionsDiscounts()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, cdList[0].ScDiscountPct.Float64, 0.2)

	err = DeleteSalesCommissionsDiscount(int32(cdId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
