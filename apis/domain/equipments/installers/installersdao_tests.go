package installers

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/ravayak/copee_backend/apis/domain/companies"
	"github.com/ravayak/copee_backend/apis/domain/equipments/products"
	"github.com/ravayak/copee_backend/apis/domain/geo"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

func TestInstallersDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Insert product to fulfill foreign key constraint
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

	// Insert company to fulfill foreign key constraint
	// Create a geo
	resInsertGeo, err := geo.InsertGeo(dt.GeoCreateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	geoId, err := resInsertGeo.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Create a company
	companyCreateParams := dt.GetCompanyCreateParams(int32(geoId))

	resInsertCompany, err := companies.InsertCompany(companyCreateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	companyId, err := resInsertCompany.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	createEqmtInstallerParams := mysql.CreateEquipmentInstallerParams{
		EiProductID: sql.NullInt32{Int32: int32(productId), Valid: true},
		EiCompanyID: sql.NullInt32{Int32: int32(companyId), Valid: true},
	}

	resInsertEqmtInstaller, err := InsertEquipmentInstaller(createEqmtInstallerParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	installerId, err := resInsertEqmtInstaller.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	installer, err := GetEquipmentInstaller(int32(installerId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, installer.EiID, int32(installerId))

	updateEqmtInstallerParams := mysql.UpdateEquipmentInstallerParams{
		EiProductID: sql.NullInt32{Int32: int32(productId), Valid: true},
		EiCompanyID: sql.NullInt32{Int32: int32(companyId), Valid: true},
		EiID:        int32(installerId),
	}

	err = UpdateEquipmentInstaller(updateEqmtInstallerParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
