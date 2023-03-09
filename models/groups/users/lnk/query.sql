-- name: GetGroupUserLnk :one
SELECT * FROM groups_users_lnk
WHERE gu_id = ? LIMIT 1;

-- name: ListGroupsUsersLnk :many
SELECT * FROM groups_users_lnk
ORDER BY gu_id ASC;

-- name: CreateGroupUserLnk :execresult
INSERT INTO groups_users_lnk (gu_user_id , gu_group_id
) VALUES ( ? , ?);

-- name: CreateGroupUserLnkByGroupName :execresult
INSERT INTO groups_users_lnk (gu_group_id, gu_user_id) SELECT `groups`.g_id, users.user_id FROM `groups` JOIN users ON users.user_id = ? WHERE `groups`.g_name = ?;
-- name: DeleteGroupUserLnk :exec
DELETE FROM groups_users_lnk
WHERE gu_id = ?;

-- name: UpdateGroupUserLnk :exec
UPDATE groups_users_lnk SET gu_group_id = ? , gu_user_id = ?
WHERE gu_id = ?;