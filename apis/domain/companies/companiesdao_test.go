package companies

import (
	"database/sql"
	"fmt"
	"testing"

	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCompaniesDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Create a company

	resInsertCompany, err := InsertCompany(dt.CreateCompanyParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	companyId, err := resInsertCompany.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get the company
	company, err := GetCompany(int32(companyId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, company.CID, int32(companyId))

	// Update the company
	companyUpdateParams := mysql.UpdateCompanyParams{
		CName:         sql.NullString{String: "Company 2", Valid: true},
		CRcs:          sql.NullString{String: "Type 2", Valid: true},
		CSiret:        sql.NullString{String: "Siret 2", Valid: true},
		CSiren:        sql.NullString{String: "Siren 2", Valid: true},
		CIntraEuVat:   sql.NullString{String: "IntraEuVat 2", Valid: true},
		CLegalStatus:  sql.NullString{String: "LegalStatus 2", Valid: true},
		CCreationDate: sql.NullString{String: "Updated", Valid: true},
		CCapital:      sql.NullInt32{Int32: 1000, Valid: true},
		CID:           int32(companyId),
	}

	// Update the company
	err = UpdateCompany(companyUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated company
	companys, err := ListCompanies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, companys[0].CCreationDate.String, "Updated")

	// Delete the company
	err = DeleteCompany(int32(companyId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
