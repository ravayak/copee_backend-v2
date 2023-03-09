package commissions

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/groups"
	"github.com/ravayak/copee_backend/apis/domain/sales/commissions/discounts"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCommissionsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Create group to fulfill FK
	CreateGroupParams := mysql.CreateGroupParams{
		GName:        sql.NullString{String: "Test", Valid: true},
		GDescription: sql.NullString{String: "Test", Valid: true},
	}

	resInsertGroup, err := groups.InsertGroup(CreateGroupParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupId, err := resInsertGroup.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Create commissions discounts to fulfill FK
	createCdParams := mysql.CreateSalesCommissionDiscountParams{
		ScDiscountPct: sql.NullFloat64{Float64: 0.1, Valid: true},
		ScComPct:      sql.NullFloat64{Float64: 0.1, Valid: true},
	}

	cd, err := discounts.InsertSalesCommissionsDiscount(createCdParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	cdId, err := cd.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	createCdParams2 := mysql.CreateSalesCommissionDiscountParams{
		ScDiscountPct: sql.NullFloat64{Float64: 0.2, Valid: true},
		ScComPct:      sql.NullFloat64{Float64: 0.2, Valid: true},
	}

	cd2, err := discounts.InsertSalesCommissionsDiscount(createCdParams2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	cd2Id, err := cd2.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	/****************************************************************/

	createCommissionsParams := mysql.CreateSalesCommissionParams{
		ScCommissionsDiscountID: sql.NullInt32{Int32: int32(cdId), Valid: true},
		ScGroupID:               sql.NullInt32{Int32: int32(groupId), Valid: true},
	}

	resCom, err := InsertSalesCommission(createCommissionsParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	comId, err := resCom.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	com, err := GetSalesCommission(int32(comId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, com.ScID, int32(comId))

	updateCommissionsParams := mysql.UpdateSalesCommissionParams{
		ScCommissionsDiscountID: sql.NullInt32{Int32: int32(cd2Id), Valid: true},
		ScGroupID:               sql.NullInt32{Int32: int32(groupId), Valid: true},
		ScID:                    int32(comId),
	}

	err = UpdateSalesCommission(updateCommissionsParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	coms, err := ListSalesCommissions()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, coms[0].ScCommissionsDiscountID.Int32, int32(cd2Id))

	err = DeleteSalesCommission(int32(comId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = discounts.DeleteSalesCommissionsDiscount(int32(cdId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = discounts.DeleteSalesCommissionsDiscount(int32(cd2Id))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = groups.DeleteGroup(int32(groupId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
