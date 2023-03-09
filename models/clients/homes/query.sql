-- name: GetClientHomeParams :one
SELECT * FROM clients_homes
WHERE ch_id = ? LIMIT 1;

-- name: ListClientsHomes :many
SELECT * FROM clients_homes
ORDER BY ch_id ASC;

-- name: CreateClientHome :execresult
INSERT INTO clients_homes (ch_client_id, ch_geo_id, ch_construction_year, ch_area, ch_residents, ch_roof_positionning, ch_roof_slope, ch_label, ch_tr, ch_huc, ch_isolation
) VALUES ( 
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: DeleteClientHome :exec
DELETE FROM clients_homes
WHERE ch_id = ?;

-- name: UpdateClientHome :exec
UPDATE clients_homes SET ch_client_id = ?, ch_geo_id = ?, ch_construction_year = ?, ch_area = ?, ch_residents = ?, ch_roof_positionning = ?, ch_roof_slope = ?, ch_label = ?, ch_tr = ?, ch_huc = ?, ch_isolation = ?
WHERE ch_id = ?;