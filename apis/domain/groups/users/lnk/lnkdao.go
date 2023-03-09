package users_lnk

import (
	"context"
	"database/sql"
	"time"

	"github.com/ravayak/copee_backend/apis/domain/groups"
	"github.com/ravayak/copee_backend/apis/domain/users"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListGroupsUsersLnk() ([]mysql.GroupsUsersLnk, error) {
	ctx := context.Background()
	groupsUsersLnk, err := mysql.QueriesDb.ListGroupsUsersLnk(ctx)
	return groupsUsersLnk, err
}

func GetGroupUserLnk(guId int32) (mysql.GroupsUsersLnk, error) {
	ctx := context.Background()
	groupUserLnk, err := mysql.QueriesDb.GetGroupUserLnk(ctx, guId)
	return groupUserLnk, err
}

func InsertGroupUserLnkByGroupName(params mysql.CreateGroupUserLnkByGroupNameParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateGroupUserLnkByGroupName(ctx, params)
	return res, err
}

func InsertGroupUsersLnk(groupUserLnkParams mysql.CreateGroupUserLnkParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateGroupUserLnk(ctx, groupUserLnkParams)
	return res, err
}

func UpdateGroupUsersLnk(groupUserLnkParams mysql.UpdateGroupUserLnkParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateGroupUserLnk(ctx, groupUserLnkParams)
	return err
}

func DeleteGroupsUsersLnk(guId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteGroupUserLnk(ctx, guId)
	return err
}

// UTILITY FUNCTIONS FOR TESTS
type GroupsDependencies struct {
	User           *sql.Result
	Group          *sql.Result
	GroupUserLnk   *sql.Result
	UserId         int64
	GroupId        int64
	GroupUserLnkId int64
}

func GetGroupsDependencies() (*GroupsDependencies, error) {
	userParams := mysql.CreateUserParams{
		UserFirstname:       sql.NullString{String: "Test", Valid: true},
		UserLastname:        sql.NullString{String: "Test", Valid: true},
		UserEmail:           sql.NullString{String: "test@test.test", Valid: true},
		UserPhone:           sql.NullString{String: "", Valid: false},
		UserIsActive:        sql.NullBool{Bool: true, Valid: true},
		UserDateCreated:     time.Now(),
		UserRecruitmentDate: sql.NullTime{Valid: false},
	}
	resInsertUser, err := users.InsertUser(userParams)
	if err != nil {
		return nil, err
	}

	// Create group to fulfill FK
	CreateGroupParams := mysql.CreateGroupParams{
		GName:        sql.NullString{String: "Test group", Valid: true},
		GDescription: sql.NullString{String: "Test Description", Valid: true},
	}

	resInsertGroup, err := groups.InsertGroup(CreateGroupParams)
	if err != nil {
		return nil, err
	}

	userId, err := resInsertUser.LastInsertId()
	if err != nil {
		return nil, err
	}

	groupId, err := resInsertGroup.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Create group user link
	groupUserLnkCreateParams := mysql.CreateGroupUserLnkParams{
		GuUserID:  sql.NullInt32{Int32: int32(userId), Valid: true},
		GuGroupID: sql.NullInt32{Int32: int32(groupId), Valid: true},
	}

	resInsertGroupUserLnk, err := InsertGroupUsersLnk(groupUserLnkCreateParams)
	if err != nil {
		return nil, err
	}

	groupUserLnkId, err := resInsertGroupUserLnk.LastInsertId()
	if err != nil {
		return nil, err
	}

	groupsDependencies := GroupsDependencies{
		User:           &resInsertUser,
		Group:          &resInsertGroup,
		GroupUserLnk:   &resInsertGroupUserLnk,
		UserId:         userId,
		GroupId:        groupId,
		GroupUserLnkId: groupUserLnkId,
	}
	return &groupsDependencies, nil
}

func DeleteGroupDepencies(groupDependencies *GroupsDependencies) error {
	err := users.DeleteUser(int32(groupDependencies.UserId))
	if err != nil {
		return err
	}

	err = groups.DeleteGroup(int32(groupDependencies.GroupId))
	if err != nil {
		return err
	}

	err = DeleteGroupsUsersLnk(int32(groupDependencies.GroupUserLnkId))
	if err != nil {
		return err
	}

	return nil
}
