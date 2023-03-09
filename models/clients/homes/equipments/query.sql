-- name: GetClientHomeEquipmentParams :one
SELECT * FROM clients_homes_equipments
WHERE che_id = ? LIMIT 1;

-- name: ListClientsHomesEquipments :many
SELECT * FROM clients_homes_equipments
ORDER BY che_id ASC;

-- name: CreateClientHomeEquipment :execresult
INSERT INTO clients_homes_equipments (che_home_id,  che_equipment_type, che_description
) VALUES ( 
   ?, ?, ?
);

-- name: DeleteClientHomeEquipment :exec
DELETE FROM clients_homes_equipments
WHERE che_id = ?;

-- name: UpdateClientHomeEquipment :exec
UPDATE clients_homes_equipments SET che_home_id=?, che_equipment_type=?, che_description=?
WHERE che_id = ?;