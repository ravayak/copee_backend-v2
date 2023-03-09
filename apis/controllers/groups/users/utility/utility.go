package groups_users_controller

import (
	"database/sql"

	lnk "github.com/ravayak/copee_backend/apis/domain/groups/users/lnk"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

func SetParamsToMysqlGroupUserParams(params *lnk.CreateGroupUserLnkParams) mysql.CreateGroupUserLnkParams {
	return mysql.CreateGroupUserLnkParams{
		GuUserID:  sql.NullInt32{Int32: params.GuUserID, Valid: true},
		GuGroupID: sql.NullInt32{Int32: params.GuGroupID, Valid: true},
	}
}
