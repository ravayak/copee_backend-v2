-- name: GetClient :one
SELECT * FROM clients
WHERE client_id = ? LIMIT 1;

-- name: ListClients :many
SELECT * FROM clients
ORDER BY client_id ASC;

-- name: ListClientsByUser :many
SELECT c.*
FROM `clients` c
INNER JOIN `clients_installations` ci ON ci.ci_client_id = c.client_id
WHERE ci.ci_user_id = ?;


-- name: ListClientsIdByUser :many
SELECT c.client_id
FROM `clients` c
INNER JOIN `clients_installations` ci ON ci.ci_client_id = c.client_id
WHERE ci.ci_user_id = ?;

-- name: CreateClient :execresult
INSERT INTO clients (client_last_name, client_first_name, client_email, client_phone, client_fiscal_year_income
) VALUES ( 
  ?, ?, ?, ?, ?
);

-- name: DeleteClient :exec
DELETE FROM clients
WHERE client_id = ?;

-- name: UpdateClient :exec
UPDATE clients SET client_last_name = ?, client_first_name= ?, client_email = ?, client_phone = ?, client_fiscal_year_income = ? 
WHERE client_id = ?;