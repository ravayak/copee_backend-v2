-- name: GetEquipmentCatalogue :one
SELECT * FROM equipments_catalogue
WHERE ec_id = ? LIMIT 1;

-- name: ListEquipmentCatalogue :many
SELECT * FROM equipments_catalogue
ORDER BY ec_id ASC;

-- name: CreateEquipmentCatalogue :execresult
INSERT INTO equipments_catalogue (ec_product_id, ec_quantity, ec_description, ec_date_added, ec_date_updated)
VALUES (?, ?, ?, ?, ?);

-- name: DeleteEquipmentCatalogue :exec
DELETE FROM equipments_catalogue WHERE ec_id = ?;

-- name: UpdateEquipmentCatalogue :exec
UPDATE equipments_catalogue SET ec_product_id = ?, ec_quantity = ?, ec_description = ?, ec_date_added = ?, ec_date_updated = ?
WHERE ec_id = ?;