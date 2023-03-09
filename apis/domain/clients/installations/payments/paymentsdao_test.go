package payments

import (
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/clients/installations"
	itd "github.com/ravayak/copee_backend/apis/domain/clients/installations/tests_dependencies"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestClientInstallationPayment(t *testing.T) {
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

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

	clientInstallationId, err := resInsertClientInstallation.LastInsertId()

	// Create installation Payments
	resInsertClientInstallationPayment, err := InsertClientInstallationPayment(dt.GetClientInstallationPaymentCreateParams(int32(clientInstallationId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientInstallationPaymentId, err := resInsertClientInstallationPayment.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	clientInstallationPayment, err := GetClientInstallationPayment(int32(clientInstallationPaymentId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientInstallationPayment.CipID, int32(clientInstallationPaymentId))

	// Update client installation payment
	err = UpdateClientInstallationPayment(dt.GetClientInstallationPaymentUpdateParams(int32(clientInstallationPaymentId), int32(clientInstallationId)))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Installation Payment List
	clientInstallationPayments, err := ListClientInstallationPayments()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, clientInstallationPayments[0].CipType.String, string(("Test updated")))

	err = DeleteClientInstallationPayment(int32(clientInstallationPaymentId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Delete Client Installation
	err = installations.DeleteClientInstallation(int32(clientInstallationId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = itd.DeleteInstallationDependencies(int32(userId), int32(sharedUserId), int32(clientId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
