-- name: GetCategory :one
SELECT * FROM equipments_categories
WHERE ec_id = ? LIMIT 1;

-- name: ListCategories :many
SELECT * FROM equipments_categories
ORDER BY ec_id ASC;

-- name: CreateCategory :execresult
INSERT INTO equipments_categories (ec_category_code, ec_description)
VALUES (?, ?);

-- name: DeleteCategory :exec
DELETE FROM equipments_categories WHERE ec_id = ?;

-- name: UpdateCategory :exec
UPDATE equipments_categories SET ec_category_code=?, ec_description=? WHERE ec_id=?