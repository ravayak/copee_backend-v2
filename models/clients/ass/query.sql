-- name: GetClientAss :one
SELECT * FROM clients_ass
WHERE ca_id = ? LIMIT 1;

-- name: ListClientsAss :many
SELECT * FROM clients_ass
ORDER BY ca_id ASC;

-- name: CreateClientAss :execresult
INSERT INTO clients_ass (ca_client_id, ca_company_id, ca_call_date, ca_call_reason, ca_intervention_date, ca_comment, ca_is_resolved) VALUES ( 
  ?, ?, ?, ?, ?, ?, ?
);

-- name: DeleteClientAss :exec
DELETE FROM clients_ass
WHERE ca_id = ?;

-- name: UpdateClientAss :exec
UPDATE clients_ass SET ca_client_id = ?, ca_company_id= ?, ca_call_date = ?, ca_call_reason = ?, ca_intervention_date = ?, ca_comment = ?, ca_is_resolved = ?
WHERE ca_id = ?;