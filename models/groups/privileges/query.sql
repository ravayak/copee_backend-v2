-- name: GetGroupPrivilege :one
SELECT * FROM groups_privileges
WHERE gp_privilege_id = ? LIMIT 1;

-- name: GetGroupPrivilegesByUserId :many
SELECT DISTINCT p.* FROM groups_privileges p
INNER JOIN groups_privileges_lnk gp ON p.gp_privilege_id = gp.gpl_privilege_id
INNER JOIN `groups` g ON gp.gpl_group_id = g.g_id
INNER JOIN groups_users_lnk gu ON g.g_id = gu.gu_group_id 
INNER JOIN users u ON gu.gu_user_id = u.user_id
WHERE u.user_id = ? ORDER BY p.gp_privilege_id ASC;

-- name: ListGroupsPrivileges :many
SELECT * FROM groups_privileges
ORDER BY gp_privilege_id ASC;

-- name: CreateGroupPrivilege :execresult
INSERT INTO groups_privileges (gp_privilege
) VALUES (?);

-- name: DeleteGroupPrivilege :exec
DELETE FROM groups_privileges
WHERE gp_privilege_id = ?;

-- name: UpdateGroupPrivilege :exec
UPDATE groups_privileges SET gp_privilege = ? 
WHERE gp_privilege_id = ?;