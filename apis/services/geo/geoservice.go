package geo

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/geo"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	GeoService geoServiceInterface = &geoService{}
)

type geoService struct{}

type geoServiceInterface interface {
	InsertGeo(geoParams mysql.CreateGeoParams) (sql.Result, error)
	UpdateGeo(geoParams mysql.UpdateGeoParams) error
	DeleteGeo(geoId int32) error
	GetGeo(geoId int32) (mysql.Geo, error)
	ListGeo() ([]mysql.Geo, error)
}

func (cs *geoService) InsertGeo(geoParams mysql.CreateGeoParams) (sql.Result, error) {
	return geo.InsertGeo(geoParams)
}

func (cs *geoService) UpdateGeo(geoParams mysql.UpdateGeoParams) error {
	return geo.UpdateGeo(geoParams)
}

func (cs *geoService) DeleteGeo(geoId int32) error {
	return geo.DeleteGeo(geoId)
}

func (cs *geoService) GetGeo(geoId int32) (mysql.Geo, error) {
	return geo.GetGeo(geoId)
}

func (cs *geoService) ListGeo() ([]mysql.Geo, error) {
	return geo.ListGeo()
}
