-- name: GetClientInstallationProduct :one
SELECT * FROM clients_installations_products
WHERE cip_id = ? LIMIT 1;

-- name: ListClientInstallationProducts :many
SELECT * FROM clients_installations_products
ORDER BY cip_id ASC;

-- name: CreateClientInstallationProduct :execresult
INSERT INTO clients_installations_products (cip_installation_id, cip_product_id, cip_discount) VALUES ( 
   ?, ?, ?
);

-- name: UpdateClientInstallationProduct :exec
UPDATE clients_installations_products SET  cip_installation_id = ?, cip_product_id = ?, cip_discount = ?
WHERE cip_id = ?;

-- name: DeleteClientInstallationProduct :exec
DELETE FROM clients_installations_products
WHERE cip_id = ?;