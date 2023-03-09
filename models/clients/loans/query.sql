-- name: GetClientLoan :one
SELECT * FROM clients_loans
WHERE cl_id = ? LIMIT 1;

-- name: ListClientsLoans :many
SELECT * FROM clients_loans
ORDER BY cl_id ASC;

-- name: CreateClientLoan :execresult
INSERT INTO clients_loans ( cl_client_id , cl_amount , cl_installments , cl_client_provision , cl_client_prepayment , cl_insured , cl_funding_agency
) VALUES ( 
  ?, ?, ?, ?, ? , ? , ?
);

-- name: DeleteClientLoan :exec
DELETE FROM clients_loans
WHERE cl_id = ?;

-- name: UpdateClientLoan :exec
UPDATE clients_loans SET cl_client_id = ?, cl_amount= ?, cl_installments = ?, cl_client_provision = ?, cl_client_prepayment = ? , cl_insured = ? , cl_funding_agency = ?
WHERE cl_id = ?;