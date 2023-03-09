-- name: GetEquipmentDesignation :one
SELECT * FROM equipments_designations
WHERE ea_id = ? LIMIT 1;

-- name: ListEquipmentDesignations :many
SELECT * FROM equipments_designations
ORDER BY ea_id ASC;

-- name: CreateEquipmentDesignations :execresult
INSERT INTO equipments_designations (ea_product_id, ea_article, ea_description, ea_title)
VALUES (?, ?, ?, ?);

-- name: UpdateEquipmentDesignation :exec
UPDATE equipments_designations SET ea_product_id = ?, ea_article = ?, ea_description = ?, ea_title = ? WHERE ea_id = ?;

-- name: DeleteEquipmentDesignation :exec
DELETE FROM equipments_designations WHERE ea_id = ?;