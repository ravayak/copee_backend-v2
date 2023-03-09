package groups

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListGroups() ([]mysql.Group, error) {
	ctx := context.Background()
	groups, err := mysql.QueriesDb.ListGroups(ctx)
	return groups, err
}

func ListGroupsByUser(userId int32) ([]mysql.Group, error) {
	ctx := context.Background()
	groups, err := mysql.QueriesDb.ListGroupsByUser(ctx, sql.NullInt32{Int32: userId, Valid: true})
	return groups, err
}

func GetGroup(groupId int32) (mysql.Group, error) {
	ctx := context.Background()
	group, err := mysql.QueriesDb.GetGroup(ctx, groupId)
	return group, err
}

func GetGroupByName(groupName string) (mysql.Group, error) {
	ctx := context.Background()
	group, err := mysql.QueriesDb.GetGroupByName(ctx, sql.NullString{String: groupName, Valid: true})
	return group, err
}

func InsertGroup(groupParams mysql.CreateGroupParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateGroup(ctx, groupParams)
	return res, err
}

func UpdateGroup(groupParams mysql.UpdateGroupParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateGroup(ctx, groupParams)
	return err
}

func DeleteGroup(groupId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteGroup(ctx, groupId)
	return err
}
