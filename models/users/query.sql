-- name: GetUser :one
SELECT * FROM users
WHERE user_id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE user_email = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id ASC;

-- name: CreateUser :execresult
INSERT INTO users (user_firstname, user_lastname, user_email, user_phone, user_is_active,  user_date_created, user_recruitment_date 
) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?;

-- name: UpdateUser :exec
UPDATE users SET user_firstname = ?, user_lastname = ?, user_email = ?, user_phone = ?, user_is_active = ? , user_date_created = ?, user_recruitment_date = ? WHERE user_id = ?;