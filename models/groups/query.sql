-- name: GetGroup :one
SELECT * FROM `groups`
WHERE g_id = ? LIMIT 1;

-- name: GetGroupByName :one
SELECT * FROM `groups`
WHERE g_name = ? LIMIT 1;


-- name: ListGroups :many
SELECT * FROM `groups`
ORDER BY g_id ASC;

-- name: ListGroupsByUser :many
SELECT g.* FROM `groups` g
INNER JOIN `groups_users_lnk` gul ON gul.gu_group_id = g.g_id
WHERE gul.gu_user_id = ? ORDER BY g.g_id ASC;

-- name: CreateGroup :execresult
INSERT INTO `groups` (g_name, g_description
) VALUES ( 
  ?, ?
);

-- name: UpdateGroup :exec
UPDATE `groups` SET g_name = ?, g_description = ?
WHERE g_id = ?;

-- name: DeleteGroup :exec
DELETE FROM `groups`
WHERE g_id = ?;