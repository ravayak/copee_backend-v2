package clients

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/ravayak/copee_backend/apis/domain/clients"
	"github.com/ravayak/copee_backend/apis/domain/companies"
	"github.com/ravayak/copee_backend/apis/domain/geo"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestClientsAssDao(t *testing.T) {
	// Initialisation de mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Création d'un client pour la Foreign Key de ClientAss
	CreateClientParams := mysql.CreateClientParams{
		ClientLastName:         sql.NullString{String: "Test", Valid: true},
		ClientFirstName:        sql.NullString{String: "Test", Valid: true},
		ClientEmail:            sql.NullString{String: "test@test.test", Valid: true},
		ClientFiscalYearIncome: sql.NullInt32{Int32: 100000, Valid: true},
	}

	resInsertClient, err := clients.InsertClient(CreateClientParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientId, err := resInsertClient.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Création d'un geo pour la Foreign Key de Company
	geoCreateParams := dt.GeoCreateParams

	resInsertGeo, err := geo.InsertGeo(geoCreateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	geoId, err := resInsertGeo.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Créaton d'une company pour la Foreign Key de ClientAss
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

	// Création d'un ClientAss
	clientAssCreateParams := mysql.CreateClientAssParams{
		CaClientID:         sql.NullInt32{Int32: int32(clientId), Valid: true},
		CaCompanyID:        sql.NullInt32{Int32: int32(companyId), Valid: true},
		CaCallDate:         sql.NullTime{Time: time.Now(), Valid: true},
		CaCallReason:       sql.NullString{String: "Test", Valid: true},
		CaInterventionDate: sql.NullTime{Valid: false},
		CaComment:          sql.NullString{String: "Test", Valid: true},
		CaIsResolved:       sql.NullBool{Bool: false, Valid: true},
	}

	resInsertClientAss, err := InsertClientAss(clientAssCreateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientAssId, err := resInsertClientAss.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// List ClientAss
	clientAss, err := GetClientAss(int32(clientAssId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientAss.CaID, int32(clientAssId))

	clientAssUpdateParams := mysql.UpdateClientAssParams{
		CaID:               int32(clientAssId),
		CaClientID:         sql.NullInt32{Int32: int32(clientId), Valid: true},
		CaCompanyID:        sql.NullInt32{Int32: int32(companyId), Valid: true},
		CaCallDate:         sql.NullTime{Time: time.Now(), Valid: true},
		CaCallReason:       sql.NullString{String: "Test", Valid: true},
		CaInterventionDate: sql.NullTime{Valid: false},
		CaComment:          sql.NullString{String: "Updated", Valid: true},
		CaIsResolved:       sql.NullBool{Bool: false, Valid: true},
	}

	// Update ClientAss
	err = UpdateClientAss(clientAssUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated ClientAss
	clientAssList, err := ListClientsAss()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientAssList[0].CaComment.String, "Updated")

	// Delete ClientAss
	err = DeleteClientAss(int32(clientAssId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Delete Company
	err = companies.DeleteCompany(int32(companyId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Delete Geo
	err = geo.DeleteGeo(int32(geoId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Delete Client
	err = clients.DeleteClient(int32(clientId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
