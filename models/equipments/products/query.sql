-- name: GetEquipmentProduct :one
SELECT * FROM equipments_products
WHERE ep_id = ? LIMIT 1;

-- name: ListEquipmentProducts :many
SELECT * FROM equipments_products
ORDER BY ep_id ASC;

-- name: CreateEquipmentProduct :execresult
INSERT INTO equipments_products (ep_product_id, ep_price, ep_vat, ep_installation_price)
VALUES ( ?, ?, ?, ?);

-- name: DeleteEquipmentProduct :exec
DELETE FROM equipments_products WHERE ep_id = ?;

-- name: UpdateEquipmentProduct :exec
UPDATE equipments_products SET ep_product_id = ?, ep_price = ?, ep_vat = ?, ep_installation_price = ?
WHERE ep_id = ?;