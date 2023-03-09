-- name: GetLog :one
SELECT * FROM logs
WHERE log_id = ? LIMIT 1;

-- name: ListLogs :many
SELECT * FROM logs
ORDER BY log_id ASC;

-- name: CreateLog :execresult
INSERT INTO logs ( log_event_desc, log_event_code, log_event_date) VALUES ( 
  ? , ? , ?
);

-- name: DeleteLog :exec
DELETE FROM logs
WHERE log_id = ?;

-- name: UpdateLog :exec
UPDATE logs SET log_event_desc = ?, log_event_code = ?, log_event_date = ?
WHERE log_id = ?;