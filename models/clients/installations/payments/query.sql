-- name: GetClientInstallationPayment :one
SELECT * FROM clients_installations_payments
WHERE cip_id = ? LIMIT 1;

-- name: ListClientsInstallationsPayments :many
SELECT * FROM clients_installations_payments
ORDER BY cip_id ASC;

-- name: CreateClientInstallationPayment :execresult
INSERT INTO clients_installations_payments (cip_amount , cip_type , cip_date , cip_installation_id , cip_description , cip_transaction_number , cip_vat_rate
) VALUES ( 
  ?, ?, ?, ?, ? , ? , ?
);

-- name: DeleteClientInstallationPayment :exec
DELETE FROM clients_installations_payments
WHERE cip_id = ?;

-- name: UpdateClientInstallationPayment :exec
UPDATE clients_installations_payments SET cip_amount = ?, cip_type= ?, cip_date = ?, cip_installation_id = ?, cip_description = ? , cip_transaction_number = ? , cip_vat_rate = ?
WHERE cip_id = ?;