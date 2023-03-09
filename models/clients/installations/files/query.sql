-- name: GetClientInstallationFile :one
SELECT * FROM clients_installations_files
WHERE cif_id = ? LIMIT 1;

-- name: ListClientsInstallationFiles :many
SELECT * FROM clients_installations_files
ORDER BY cif_id ASC;

-- name: CreateClientInstallationFile :execresult
INSERT INTO clients_installations_files (cif_installation_id , cif_file_type , cif_file_url , cif_file_creation_date
) VALUES ( 
  ?, ?, ?, ?
);

-- name: DeleteClientInstallationFile :exec
DELETE FROM clients_installations_files
WHERE cif_id = ?;

-- name: UpdateClientInstallationFile :exec
UPDATE clients_installations_files SET cif_installation_id = ?, cif_file_type = ?, cif_file_url = ?, cif_file_creation_date = ?
WHERE cif_id = ?;