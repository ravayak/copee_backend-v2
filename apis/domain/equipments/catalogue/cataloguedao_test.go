package catalogue

import (
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/equipments/products"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestEquipmentCatalogueDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Creating product to fulfill product FK in catalogue
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

	// Creating catalogue
	resInsertEquipmentCatalogue, err := InsertEquipmentCatalogue(dt.GetEquipmentCatalogueParams(int32(productId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	catalogueId, err := resInsertEquipmentCatalogue.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	catalogue, err := GetEquipmentCatalogue(int32(catalogueId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, catalogue.EcDescription.String, "Test")

	resUpdateEquipmentCatalogue := dt.GetUpdateEquipmentCatalogueParams(int32(catalogueId), int32(productId))

	err = UpdateEquipmentCatalogue(resUpdateEquipmentCatalogue)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	catalogues, err := ListEquipmentCatalogue()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, catalogues[0].EcDescription.String, "Test updated")

	err = DeleteEquipmentCatalogue(int32(catalogueId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
