-- name: GetClientHomeParamsBill :one
SELECT * FROM clients_homes_bills
WHERE chb_id = ? LIMIT 1;

-- name: ListClientsHomesBills :many
SELECT * FROM clients_homes_bills
ORDER BY chb_id ASC;

-- name: CreateClientHomeBill :execresult
INSERT INTO clients_homes_bills (chb_home_id, chb_electricity, chb_natural_gas , chb_propane_gas , chb_wood , chb_heating_oil, chb_year 
) VALUES ( 
  ?, ? , ? , ? , ? , ?, ?
);

-- name: DeleteClientHomeBill :exec
DELETE FROM clients_homes_bills
WHERE chb_id = ?;

-- name: UpdateClientHomeBill :exec
UPDATE clients_homes_bills SET chb_home_id=?, chb_electricity=?, chb_natural_gas=?, chb_propane_gas=?, chb_wood=?, chb_heating_oil=?, chb_year=?
WHERE chb_id = ?;