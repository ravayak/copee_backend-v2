package companiesservice

import (
	"github.com/ravayak/copee_backend/apis/domain/companies"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	CompaniesService companiesServiceInterface = &companiesService{}
)

type companiesService struct{}

type companiesServiceInterface interface {
	InsertCompany(companyParams mysql.CreateCompanyParams) (*companies.Company, error)
	UpdateCompany(companyParams mysql.UpdateCompanyParams) error
	DeleteCompany(companyId int32) error
	GetCompany(companyId int32) (*companies.Company, error)
	ListCompanies() ([]*companies.Company, error)
}

func (cs *companiesService) InsertCompany(companyParams mysql.CreateCompanyParams) (*companies.Company, error) {

	res, err := companies.InsertCompany(companyParams)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	company, err := cs.GetCompany(int32(id))
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (cs *companiesService) UpdateCompany(companyParams mysql.UpdateCompanyParams) error {
	return companies.UpdateCompany(companyParams)
}

func (cs *companiesService) DeleteCompany(companyId int32) error {
	return companies.DeleteCompany(companyId)
}

func (cs *companiesService) GetCompany(companyId int32) (*companies.Company, error) {
	companies, err := companies.GetCompany(companyId)
	if err != nil {
		return nil, err
	}

	return convertToCompanieResponse(companies), nil
}

func (cs *companiesService) ListCompanies() ([]*companies.Company, error) {
	mysqlCompanies, err := companies.ListCompanies()
	if err != nil {
		return nil, err
	}

	companiesRes := convertToCompanieListResponse(mysqlCompanies)
	return companiesRes, nil
}

//  Utility functions to convert from mysql to companies.Company in response

func convertToCompanieResponse(company mysql.Company) *companies.Company {
	return &companies.Company{
		CID:           company.CID,
		CGeoID:        company.CGeoID.Int32,
		CName:         company.CName.String,
		CRcs:          company.CRcs.String,
		CSiret:        company.CSiret.String,
		CSiren:        company.CSiren.String,
		CIntraEuVat:   company.CIntraEuVat.String,
		CLegalStatus:  company.CLegalStatus.String,
		CCreationDate: company.CCreationDate.String,
		CCapital:      company.CCapital.Int32,
	}
}

func convertToCompanieListResponse(mysqlCompanies []mysql.Company) []*companies.Company {
	companiesResponse := []*companies.Company{}
	for _, company := range mysqlCompanies {
		companiesResponse = append(companiesResponse, convertToCompanieResponse(company))
	}
	return companiesResponse
}
