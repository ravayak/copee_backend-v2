package geo

import (
	"database/sql"
	"fmt"
	"testing"

	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGeoDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Create a geo
	geoCreateParams := mysql.CreateGeoParams{
		ChgZone:       sql.NullString{String: "h2", Valid: true},
		ChgLatitude:   sql.NullFloat64{Float64: 45.5, Valid: true},
		ChgLongitude:  sql.NullFloat64{Float64: 45.5, Valid: true},
		ChgAltitude:   sql.NullInt32{Int32: 45, Valid: true},
		ChgDepartment: sql.NullInt32{Int32: 66, Valid: true},
		ChgCity:       sql.NullString{String: "Paris", Valid: true},
		ChgAddress:    sql.NullString{String: "1 rue de la paix", Valid: true},
		ChgPostcode:   sql.NullInt32{Int32: 75001, Valid: true},
	}

	resInsertGeo, err := InsertGeo(geoCreateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	geoId, err := resInsertGeo.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Geo
	geo, err := GetGeo(int32(geoId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, geo.ChgID, int32(geoId))

	geoUpdateParams := mysql.UpdateGeoParams{
		ChgZone:       sql.NullString{String: "h3", Valid: true},
		ChgLatitude:   sql.NullFloat64{Float64: 45.5, Valid: true},
		ChgLongitude:  sql.NullFloat64{Float64: 45.5, Valid: true},
		ChgAltitude:   sql.NullInt32{Int32: 45, Valid: true},
		ChgDepartment: sql.NullInt32{Int32: 66, Valid: true},
		ChgCity:       sql.NullString{String: "City Updated", Valid: true},
		ChgAddress:    sql.NullString{String: "adress Updated", Valid: true},
		ChgPostcode:   sql.NullInt32{Int32: 75001, Valid: true},
		ChgID:         int32(geoId),
	}

	// Update Geo
	err = UpdateGeo(geoUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated Geo
	geos, err := ListGeo()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, geos[0].ChgCity.String, "City Updated")

	// Delete Geo
	err = DeleteGeo(int32(geoId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
