package products

import (
	"fmt"
	"testing"

	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestProductsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	resInsertProduct, err := InsertEquipmentProduct(dt.CreateEquipmentProductParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	productId, err := resInsertProduct.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	product, err := GetEquipmentProduct(int32(productId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, product.EpID, int32(productId))

	equipmentProductUpdatedParams := dt.GetUpdateEquipmentProductParams(int32(productId))
	err = UpdateEquipmentProduct(equipmentProductUpdatedParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	products, err := ListEquipmentProducts()
	assert.Equal(t, products[0].EpPrice.String, "11.00")

	err = DeleteEquipmentProduct(int32(productId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
