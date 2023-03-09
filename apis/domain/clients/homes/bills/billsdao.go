package bills

import (
	"context"
	"database/sql"

	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	sliceUtils "github.com/ravayak/copee_backend/app/utils/slice"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	listClientHomeBillsByHomesQuery = "SELECT b.* FROM `clients_homes_bills` b INNER JOIN `clients_homes` m ON m.ch_id = b.chb_home_id WHERE m.ch_id IN "
)

func ListClientHomesBills() ([]mysql.ClientsHomesBill, error) {
	ctx := context.Background()
	client_homes_bills, err := mysql.QueriesDb.ListClientsHomesBills(ctx)
	return client_homes_bills, err
}

func ListClientHomeBillsByHomes(homeIds []int32) ([]mysql.ClientsHomesBill, error) {
	stmt, err := mysql.Db.Prepare(mysqlutils.CompleteQuery(listClientHomeBillsByHomesQuery, homeIds))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sliceUtils.ConvertInt32ArrayToInterfaceArray(homeIds)...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var homeBills []mysql.ClientsHomesBill
	for rows.Next() {
		homeBill := mysql.ClientsHomesBill{}
		err := rows.Scan(&homeBill.ChbID, &homeBill.ChbHomeID, &homeBill.ChbElectricity, &homeBill.ChbNaturalGas, &homeBill.ChbPropaneGas, &homeBill.ChbWood, &homeBill.ChbHeatingOil, &homeBill.ChbYear)
		if err != nil {
			panic(err.Error())
		}
		homeBills = append(homeBills, homeBill)
	}

	return homeBills, nil
}

func GetClientHomeParamsBill(clientId int32) (mysql.ClientsHomesBill, error) {
	ctx := context.Background()
	client_homes_bill, err := mysql.QueriesDb.GetClientHomeParamsBill(ctx, clientId)
	return client_homes_bill, err
}

func InsertClientHomesBill(clientParams mysql.CreateClientHomeBillParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientHomeBill(ctx, clientParams)
	return res, err
}

func UpdateClientHomesBill(clientParams mysql.UpdateClientHomeBillParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientHomeBill(ctx, clientParams)
	return err
}

func DeleteClientHomesBill(clientId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientHomeBill(ctx, clientId)
	return err
}
