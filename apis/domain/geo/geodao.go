package geo

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListGeo() ([]mysql.Geo, error) {
	ctx := context.Background()
	geo, err := mysql.QueriesDb.ListGeo(ctx)
	return geo, err
}

func GetGeo(geoId int32) (mysql.Geo, error) {
	ctx := context.Background()
	geo, err := mysql.QueriesDb.GetGeo(ctx, geoId)
	return geo, err
}

func InsertGeo(geoParams mysql.CreateGeoParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateGeo(ctx, geoParams)
	return res, err
}

func UpdateGeo(geoParams mysql.UpdateGeoParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateGeo(ctx, geoParams)
	return err
}

func DeleteGeo(geoId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteGeo(ctx, geoId)
	return err
}
