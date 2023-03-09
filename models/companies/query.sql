-- name: GetCompany :one
SELECT * FROM companies
WHERE c_id = ? LIMIT 1;

-- name: ListCompanies :many
SELECT * FROM companies ORDER BY c_id ASC;

-- name: CreateCompany :execresult
INSERT INTO companies (c_geo_id, c_name, c_rcs, c_siret, c_siren, c_intra_eu_vat, c_legal_status, c_creation_date, c_capital) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteCompany :exec
DELETE FROM companies
WHERE c_id = ?;

-- name: UpdateCompany :exec
UPDATE companies SET c_geo_id = ?, c_name = ?, c_rcs = ?, c_siret = ?, c_siren = ?, c_intra_eu_vat = ?, c_legal_status = ?, c_creation_date = ?, c_capital = ?
WHERE c_id = ?;