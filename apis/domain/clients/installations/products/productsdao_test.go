package products

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/clients/installations"
	itd "github.com/ravayak/copee_backend/apis/domain/clients/installations/tests_dependencies"
	"github.com/ravayak/copee_backend/apis/domain/equipments/products"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestEquipmentProductsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Creating installation & product to fulfill FK
	clientId, userId, sharedUserId, err := itd.CreateInstallationDepencencies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resInsertClientInstallation, err := installations.InsertClientInstallation(dt.GetClientInstallationCreateParams(int32(clientId), int32(userId), int32(sharedUserId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	installationId, err := resInsertClientInstallation.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

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

	createProductsParams := mysql.CreateClientInstallationProductParams{
		CipInstallationID: sql.NullInt32{Int32: int32(installationId), Valid: true},
		CipProductID:      sql.NullInt32{Int32: int32(productId), Valid: true},
		CipDiscount:       sql.NullString{String: "10.00", Valid: true},
	}

	resInsertClientInstallationProduct, err := InsertClientInstallationProduct(createProductsParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	installationProductId, err := resInsertClientInstallationProduct.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	installationProduct, err := GetClientInstallationProduct(int32(installationProductId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, installationProduct.CipDiscount.String, "10.00")

	resUpdateInstallationProduct := mysql.UpdateClientInstallationProductParams{
		CipID:             int32(installationProductId),
		CipInstallationID: sql.NullInt32{Int32: int32(installationId), Valid: true},
		CipProductID:      sql.NullInt32{Int32: int32(productId), Valid: true},
		CipDiscount:       sql.NullString{String: "11.00", Valid: true},
	}

	err = UpdateClientInstallationProduct(resUpdateInstallationProduct)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	products, err := ListClientInstallationProducts()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, products[0].CipDiscount.String, "11.00")

	err = DeleteClientInstallationProduct(int32(installationProductId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	itd.DeleteInstallationDependencies(int32(clientId), int32(userId), int32(sharedUserId))
}
