-- name: GetClientInstallation :one
SELECT * FROM clients_installations
WHERE ci_id = ? LIMIT 1;

-- name: ListClientsInstallation :many
SELECT * FROM clients_installations
ORDER BY ci_id ASC;

-- name: CreateClientInstallation :execresult
INSERT INTO clients_installations (ci_client_id , ci_user_id , ci_shared_user_id , ci_creation_date , ci_update_date , ci_status, ci_comments) VALUES ( 
   ?, ?, ?, ?, ?, ?, ?
);

-- name: DeleteClientInstallation :exec
DELETE FROM clients_installations
WHERE ci_id = ?;

-- name: UpdateClientInstallation :exec
UPDATE clients_installations SET ci_client_id=?, ci_user_id=?, ci_shared_user_id=?, ci_creation_date=?, ci_update_date=?, ci_status=?, ci_comments=?
WHERE ci_id = ?;