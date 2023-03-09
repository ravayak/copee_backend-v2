-- name: GetGroupPrivilegeLnk :one
SELECT * FROM groups_privileges_lnk
WHERE gpl_id = ? LIMIT 1;

-- name: ListGroupPrivilegesLnk :many
SELECT * FROM groups_privileges_lnk
ORDER BY gpl_id ASC;

-- name: CreateGroupPrivilegeLnk :execresult
INSERT INTO groups_privileges_lnk (gpl_group_id , gpl_privilege_id
) VALUES ( 
  ? , ?
);

-- name: UpdateGroupPrivilegeLnk :exec
UPDATE groups_privileges_lnk SET gpl_group_id = ? , gpl_privilege_id = ?
WHERE gpl_id = ?;

-- name: DeleteGroupPrivilegeLnk :exec
DELETE FROM groups_privileges_lnk
WHERE gpl_id = ?;
