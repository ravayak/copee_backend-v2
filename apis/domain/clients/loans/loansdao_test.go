package loans

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/clients"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestClientsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")

	// Create clients to fullfil Foreign Key constraint
	resInsertClient, err := clients.InsertClient(dt.CreateClientParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientId, err := resInsertClient.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	resInsertClientLoan, err := InsertClientLoan(dt.GetClientLoanCreateParams(int32(clientId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientLoanId, err := resInsertClientLoan.LastInsertId()

	// Get Client
	clientLoan, err := GetClientLoan(int32(clientLoanId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientLoan.ClClientID.Int32, int32(clientId))
	assert.Equal(t, clientLoan.ClID, int32(clientLoanId))

	clientUpdateParams := mysql.UpdateClientLoanParams{
		ClFundingAgency:    sql.NullString{String: "Test Updated", Valid: true},
		ClAmount:           clientLoan.ClAmount,
		ClInstallments:     clientLoan.ClInstallments,
		ClClientProvision:  clientLoan.ClClientProvision,
		ClClientPrepayment: clientLoan.ClClientPrepayment,
		ClInsured:          clientLoan.ClInsured,
		ClID:               int32(clientLoanId),
	}

	// Update Client
	err = UpdateClientLoan(clientUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated Client
	clientsLoans, err := ListClientsLoans()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, clientsLoans[0].ClFundingAgency.String, "Test Updated")

	// Delete Client
	err = clients.DeleteClient(int32(clientId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Delete Client Loan
	err = DeleteClientLoan(int32(clientLoanId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
