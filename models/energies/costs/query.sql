-- name: GetEnergiesCost :one
SELECT * FROM energies_costs
WHERE enc_id = ? LIMIT 1;

-- name: ListEnergiesCosts :many
SELECT * FROM energies_costs
ORDER BY enc_id ASC;

-- name: CreateEnergiesCost :execresult
INSERT INTO energies_costs (enc_type, enc_date, enc_cost)
VALUES (?, ?, ?);

-- name: DeleteEnergiesCost :exec
DELETE FROM energies_costs WHERE enc_id = ?;

-- name: UpdateEnergiesCost :exec
UPDATE energies_costs SET enc_type = ?, enc_date = ?, enc_cost = ? WHERE enc_id = ?;