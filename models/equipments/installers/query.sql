-- name: GetEquipmentInstaller :one
SELECT * FROM equipments_installers
WHERE ei_id = ? LIMIT 1;

-- name: ListEquipmentInstallers :many
SELECT * FROM equipments_installers
ORDER BY ei_id ASC;

-- name: CreateEquipmentInstaller :execresult
INSERT INTO equipments_installers (ei_product_id, ei_company_id)
VALUES (?, ?);

-- name: DeleteEquipmentInstaller :exec
DELETE FROM equipments_installers WHERE ei_id = ?;

-- name: UpdateEquipmentInstaller :exec
UPDATE equipments_installers SET ei_product_id = ?, ei_company_id = ?
WHERE ei_id = ?;