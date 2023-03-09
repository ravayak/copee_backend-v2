package homes

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/clients"
	"github.com/ravayak/copee_backend/apis/domain/geo"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	sliceUtils "github.com/ravayak/copee_backend/app/utils/slice"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	listClientHomeByClientsQuery = "SELECT m.* FROM `clients_homes` m INNER JOIN `clients` c ON c.client_id = m.ch_client_id WHERE c.client_id IN "
)

func ListClientHome() ([]mysql.ClientsHome, error) {
	ctx := context.Background()
	client_homes, err := mysql.QueriesDb.ListClientsHomes(ctx)
	return client_homes, err
}

/*
sqlc doesnt support IN clause in queries
yet with multiple parameters like with postgresql
(see doc: https://docs.sqlc.dev/en/stable/howto/select.html )
See also this github issue:

	Support passing slices to MySQL queries : https://github.com/kyleconroy/sqlc/issues/695

Therefore we need to write the query manually
*/
func ListClientHomesByClients(clientIDs []int32) ([]mysql.ClientsHome, error) {
	stmt, err := mysql.Db.Prepare(mysqlutils.CompleteQuery(listClientHomeByClientsQuery, clientIDs))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(sliceUtils.ConvertInt32ArrayToInterfaceArray(clientIDs)...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var homes []mysql.ClientsHome
	for rows.Next() {
		home := mysql.ClientsHome{}
		err := rows.Scan(&home.ChID, &home.ChClientID, &home.ChGeoID, &home.ChConstructionYear, &home.ChArea, &home.ChResidents, &home.ChRoofPositionning, &home.ChRoofSlope, &home.ChLabel, &home.ChTr, &home.ChHuc, &home.ChIsolation)
		if err != nil {
			panic(err.Error())
		}
		homes = append(homes, home)
	}

	return homes, nil
}

func GetClientHomeParams(clientId int32) (mysql.ClientsHome, error) {
	ctx := context.Background()
	client_homes, err := mysql.QueriesDb.GetClientHomeParams(ctx, clientId)
	return client_homes, err
}

func InsertClientHome(clientParams mysql.CreateClientHomeParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateClientHome(ctx, clientParams)
	return res, err
}

func UpdateClientHome(clientParams mysql.UpdateClientHomeParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateClientHome(ctx, clientParams)
	return err
}

func DeleteClientHome(clientId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteClientHome(ctx, clientId)
	return err
}

// UTILITY FUNCTIONS FOR TESTS
func CreateClientHomesDependencies() (int64, int64, error) {
	// Client Create
	clientCreateParams := dt.CreateClientParams

	resInsertClient, err := clients.InsertClient(clientCreateParams)
	if err != nil {
		return -1, -1, nil
	}

	clientId, err := resInsertClient.LastInsertId()
	if err != nil {
		return -1, -1, nil
	}

	// Geo Create
	resInsertGeo, err := geo.InsertGeo(dt.GeoCreateParams)
	if err != nil {
		return -1, -1, nil
	}

	geoId, err := resInsertGeo.LastInsertId()
	if err != nil {
		return -1, -1, nil
	}

	return geoId, clientId, nil
}

func DeleteClientHomesDependencies(clientId, geoId int32) error {
	err := geo.DeleteGeo(geoId)
	if err != nil {
		return err
	}

	err = clients.DeleteClient(clientId)
	if err != nil {
		return err
	}

	return nil
}
