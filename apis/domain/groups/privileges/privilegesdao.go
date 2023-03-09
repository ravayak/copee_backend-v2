package privileges

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func GetGroupPrivilegesByUserId(userId int32) ([]mysql.GroupsPrivilege, error) {
	ctx := context.Background()
	privileges, err := mysql.QueriesDb.GetGroupPrivilegesByUserId(ctx, userId)
	return privileges, err
}

func ListGroupsPrivileges() ([]mysql.GroupsPrivilege, error) {
	ctx := context.Background()
	privileges, err := mysql.QueriesDb.ListGroupsPrivileges(ctx)
	return privileges, err
}

func GetGroupPrivilege(privId int32) (mysql.GroupsPrivilege, error) {
	ctx := context.Background()
	privilege, err := mysql.QueriesDb.GetGroupPrivilege(ctx, privId)
	return privilege, err
}

func InsertGroupPrivilege(privilege sql.NullString) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateGroupPrivilege(ctx, privilege)
	return res, err
}

func UpdateGroupPrivilege(privilegeParams mysql.UpdateGroupPrivilegeParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateGroupPrivilege(ctx, privilegeParams)
	return err
}

func DeleteGroupPrivilege(privId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteGroupPrivilege(ctx, privId)
	return err
}
