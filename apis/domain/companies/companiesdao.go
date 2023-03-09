package companies

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListCompanies() ([]mysql.Company, error) {
	ctx := context.Background()
	companies, err := mysql.QueriesDb.ListCompanies(ctx)
	return companies, err
}

func GetCompany(companyId int32) (mysql.Company, error) {
	ctx := context.Background()
	company, err := mysql.QueriesDb.GetCompany(ctx, companyId)
	return company, err
}

func InsertCompany(companyParams mysql.CreateCompanyParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateCompany(ctx, companyParams)
	return res, err
}

func UpdateCompany(companyParams mysql.UpdateCompanyParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateCompany(ctx, companyParams)
	return err
}

func DeleteCompany(companyId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteCompany(ctx, companyId)
	return err
}
