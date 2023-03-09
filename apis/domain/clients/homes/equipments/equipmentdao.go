package equipment

import (
	"context"
	"database/sql"

	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	sliceUtils "github.com/ravayak/copee_backend/app/utils/slice"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	listClientHomesEquipmentByHomesQuery = "SELECT e.* FROM `clients_homes_equipments` e INNER JOIN `clients_homes` m ON m.ch_id = e.che_home_id WHERE m.ch_id IN "
)

func ListClientHomeEquipments() ([]mysql.ClientsHomesEquipment, error) {
	ctx := context.Background()
	client_homes_equipment, err := mysql.QueriesDb.ListClientsHomesEquipments(ctx)
	return client_homes_equipment, err
}

func ListClientHomeEquipmentsByHomes(homeIds []int32) ([]mysql.ClientsHomesEquipment, error) {
	stmt, err := mysql.Db.Prepare(mysqlutils.CompleteQuery(listClientHomesEquipmentByHomesQuery, homeIds))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sliceUtils.ConvertInt32ArrayToInterfaceArray(homeIds)...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var homeEqmts []mysql.ClientsHomesEquipment
	for rows.Next() {
		homeEqmt := mysql.ClientsHomesEquipment{}
		err := rows.Scan(&homeEqmt.CheID, &homeEqmt.CheHomeID, &homeEqmt.CheEquipmentType, &homeEqmt.CheDescription)
		if err != nil {
			panic(err.Error())
		}
		homeEqmts = append(homeEqmts, homeEqmt)
	}

	return homeEqmts, nil
}

func GetClientHomeEquipmentParams(clientId int32) (mysql.ClientsHomesEquipment, error) {
	ctx := context.Background()
	client_homes_equipment, err := mysql.QueriesDb.GetClientHomeEquipmentParams(ctx, clientId)
	return client_homes_equipment, err
}

func InsertClientHomeEquipment(clientParams mysql.CreateClientHomeEquipmentParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientHomeEquipment(ctx, clientParams)
	return res, err
}

func UpdateClientHomeEquipment(clientParams mysql.UpdateClientHomeEquipmentParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientHomeEquipment(ctx, clientParams)
	return err
}

func DeleteClientHomeEquipment(clientId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientHomeEquipment(ctx, clientId)
	return err
}
