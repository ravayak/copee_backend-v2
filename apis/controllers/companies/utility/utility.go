package companiescontroller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	companies "github.com/ravayak/copee_backend/apis/domain/companies"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

const CREATE_COMPANIES_PARAMS = "CREATE_COMPANIES_PARAMS"
const UPDATE_COMPANIES_PARAMS = "UPDATE_COMPANIES_PARAMS"

// setCompaniesParams sets the companies params for create and update
func SetCompaniesParams(companiesParamsType string, companiesToEdit companies.Company, c *gin.Context) (interface{}, error) {
	if err := c.ShouldBindJSON(&companiesToEdit); err != nil {
		return nil, err
	}

	var companiesParams interface{}
	if companiesParamsType == CREATE_COMPANIES_PARAMS {
		companiesParams = mysql.CreateCompanyParams{
			CGeoID:        sql.NullInt32{Int32: companiesToEdit.CGeoID, Valid: true},
			CName:         sql.NullString{String: companiesToEdit.CName, Valid: true},
			CRcs:          sql.NullString{String: companiesToEdit.CRcs, Valid: true},
			CSiret:        sql.NullString{String: companiesToEdit.CSiret, Valid: true},
			CSiren:        sql.NullString{String: companiesToEdit.CSiren, Valid: true},
			CIntraEuVat:   sql.NullString{String: companiesToEdit.CIntraEuVat, Valid: true},
			CLegalStatus:  sql.NullString{String: companiesToEdit.CLegalStatus, Valid: true},
			CCreationDate: sql.NullString{String: companiesToEdit.CCreationDate, Valid: true},
			CCapital:      sql.NullInt32{Int32: companiesToEdit.CCapital, Valid: true},
		}
	}

	if companiesParamsType == UPDATE_COMPANIES_PARAMS {
		companiesParams = mysql.UpdateCompanyParams{
			CGeoID:        sql.NullInt32{Int32: companiesToEdit.CGeoID, Valid: true},
			CName:         sql.NullString{String: companiesToEdit.CName, Valid: true},
			CRcs:          sql.NullString{String: companiesToEdit.CRcs, Valid: true},
			CSiret:        sql.NullString{String: companiesToEdit.CSiret, Valid: true},
			CSiren:        sql.NullString{String: companiesToEdit.CSiren, Valid: true},
			CIntraEuVat:   sql.NullString{String: companiesToEdit.CIntraEuVat, Valid: true},
			CLegalStatus:  sql.NullString{String: companiesToEdit.CLegalStatus, Valid: true},
			CCreationDate: sql.NullString{String: companiesToEdit.CCreationDate, Valid: true},
			CCapital:      sql.NullInt32{Int32: companiesToEdit.CCapital, Valid: true},
			CID:           companiesToEdit.CID,
		}
	}

	return companiesParams, nil
}
