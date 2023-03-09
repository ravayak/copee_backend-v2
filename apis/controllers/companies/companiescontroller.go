package companiescontroller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	cpu "github.com/ravayak/copee_backend/apis/controllers/companies/utility"
	"github.com/ravayak/copee_backend/apis/domain/companies"
	cps "github.com/ravayak/copee_backend/apis/services/companies"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

func Get(c *gin.Context) {
	companyId, exists := c.Get("company_id")
	cpId := companyId.(int32)
	if exists {
		company, err := cps.CompaniesService.GetCompany(cpId)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, company)
		return
	}
	c.JSON(404, gin.H{"error": fmt.Sprintf("error Companies %d not found ", companyId)})
}

func Delete(c *gin.Context) {
	companyId, exists := c.Get("company_id")
	if exists {
		cpId := companyId.(int32)
		cps.CompaniesService.DeleteCompany(cpId)
		c.JSON(200, fmt.Sprintf("companie %d deleted ", cpId))
		return
	}
	c.JSON(404, gin.H{"error": fmt.Sprintf("error Company %d not found", companyId)})
}

func Update(c *gin.Context) {
	companyId, exists := c.Get("company_id")
	cpId := companyId.(int32)
	if exists {
		updatedCompany := companies.Company{CID: cpId}
		companieParams, err := cpu.SetCompaniesParams(cpu.UPDATE_COMPANIES_PARAMS, updatedCompany, c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = cps.CompaniesService.UpdateCompany(companieParams.(mysql.UpdateCompanyParams))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, fmt.Sprintf("companie %d updated ", cpId))
		return
	}
	c.JSON(404, gin.H{"error": fmt.Sprintf("error Company %d not found ", companyId)})
}

func Create(c *gin.Context) {
	newCompany := companies.Company{}
	companiesParams, err := cpu.SetCompaniesParams(cpu.CREATE_COMPANIES_PARAMS, newCompany, c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	company, err := cps.CompaniesService.InsertCompany(companiesParams.(mysql.CreateCompanyParams))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, company)
}

func List(c *gin.Context) {
	companies, err := cps.CompaniesService.ListCompanies()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, companies)
}
