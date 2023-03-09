package designations

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/equipments/products"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestEquipmentDesignationsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Creating product to fulfill product FK in designations
	resInsertProduct, err := products.InsertEquipmentProduct(dt.CreateEquipmentProductParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	productId, err := resInsertProduct.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	createDesignationParams := mysql.CreateEquipmentDesignationsParams{
		EaProductID:   sql.NullInt32{Int32: int32(productId), Valid: true},
		EaArticle:     sql.NullString{String: "Test", Valid: true},
		EaDescription: sql.NullString{String: "Test", Valid: true},
		EaTitle:       sql.NullString{String: "Test", Valid: true},
	}

	resInsertEquipmentDesignation, err := InsertEquipmentDesignation(createDesignationParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	designationId, err := resInsertEquipmentDesignation.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	designation, err := GetEquipmentDesignation(int32(designationId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, designation.EaDescription.String, "Test")

	resUpdateEquipmentDesignation := mysql.UpdateEquipmentDesignationParams{
		EaID:          int32(designationId),
		EaProductID:   sql.NullInt32{Int32: int32(productId), Valid: true},
		EaArticle:     sql.NullString{String: "Test Updated", Valid: true},
		EaDescription: sql.NullString{String: "Test updated", Valid: true},
		EaTitle:       sql.NullString{String: "Test Updated", Valid: true},
	}

	err = UpdateEquipmentDesignation(resUpdateEquipmentDesignation)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	designations, err := ListEquipmentDesignations()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, designations[0].EaDescription.String, "Test updated")

	err = DeleteEquipmentDesignation(int32(designationId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
