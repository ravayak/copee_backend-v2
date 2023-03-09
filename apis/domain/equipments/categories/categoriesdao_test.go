package categories

import (
	"database/sql"
	"fmt"
	"testing"

	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestEquipmentCategoriesDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	createCategoryParams := mysql.CreateCategoryParams{
		EcCategoryCode: sql.NullString{String: "Test", Valid: true},
		EcDescription:  sql.NullString{String: "Test", Valid: true},
	}

	resInsertEquipmentCategory, err := InsertEquipmentCategory(createCategoryParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	categoryId, err := resInsertEquipmentCategory.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	category, err := GetEquipmentCategory(int32(categoryId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, category.EcDescription.String, "Test")

	resUpdateEquipmentCategory := mysql.UpdateCategoryParams{
		EcID:           int32(categoryId),
		EcCategoryCode: sql.NullString{String: "Test updated", Valid: true},
		EcDescription:  sql.NullString{String: "Test updated", Valid: true},
	}

	err = UpdateEquipmentCategory(resUpdateEquipmentCategory)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	categories, err := ListEquipmentCategories()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, categories[0].EcDescription.String, "Test updated")

	err = DeleteEquipmentCategory(int32(categoryId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
