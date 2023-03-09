package privileges_lnk

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListGroupsPrivilegesLnk() ([]mysql.GroupsPrivilegesLnk, error) {
	ctx := context.Background()
	privilegesLnk, err := mysql.QueriesDb.ListGroupPrivilegesLnk(ctx)
	return privilegesLnk, err
}

func GetGroupPrivilegeLnk(privId int32) (mysql.GroupsPrivilegesLnk, error) {
	ctx := context.Background()
	privilegeLnk, err := mysql.QueriesDb.GetGroupPrivilegeLnk(ctx, privId)
	return privilegeLnk, err
}

func InsertGroupPrivilegeLnk(privilegeLnkParams mysql.CreateGroupPrivilegeLnkParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateGroupPrivilegeLnk(ctx, privilegeLnkParams)
	return res, err
}

func UpdateGroupPrivilegeLnk(privilegeParams mysql.UpdateGroupPrivilegeLnkParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateGroupPrivilegeLnk(ctx, privilegeParams)
	return err
}

func DeleteGroupPrivilegeLnk(privId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteGroupPrivilege(ctx, privId)
	return err
}
