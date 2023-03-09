-- name: GetEquipmentService :one
SELECT * FROM equipments_services
WHERE es_id = ? LIMIT 1;

-- name: ListEquipmentServices :many
SELECT * FROM equipments_services
ORDER BY es_id ASC;

-- name: CreateEquipmentServices :execresult
INSERT INTO equipments_services (es_description, es_price)
VALUES (?, ?);

-- name: DeleteEquipmentService :exec
DELETE FROM equipments_services WHERE es_id = ?;

-- name: UpdateEquipmentService :exec
UPDATE equipments_services SET es_description = ?, es_price = ?
WHERE es_id = ?;