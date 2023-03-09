-- name: GetGeo :one
SELECT * FROM geo
WHERE chg_id = ? LIMIT 1;

-- name: ListGeo :many
SELECT * FROM geo
ORDER BY chg_id ASC;

-- name: CreateGeo :execresult
INSERT INTO geo ( chg_zone , chg_latitude , chg_longitude , chg_altitude , chg_department , chg_city , chg_address , chg_postcode
) VALUES ( 
  ?, ?, ?, ?, ? , ? , ? , ?
);

-- name: DeleteGeo :exec
DELETE FROM geo
WHERE chg_id = ?;

-- name: UpdateGeo :exec
UPDATE geo SET chg_zone = ?, chg_latitude = ?, chg_longitude = ?, chg_altitude = ?, chg_department = ?, chg_city = ? , chg_address = ?, chg_postcode = ?
WHERE chg_id = ?;