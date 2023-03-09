package groups

import (
	"database/sql"

	gul "github.com/ravayak/copee_backend/apis/domain/groups/users/lnk"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	GroupsUsersService groupsUsersServiceInterface = &groupUsersService{}
)

type groupUsersService struct{}

type groupsUsersServiceInterface interface {
	InsertGroupUser(groupUserParams mysql.CreateGroupUserLnkParams) (*gul.GroupsUsersLnk, error)
	InsertGroupUserByGroupName(userId int32, groupName string) (*gul.GroupsUsersLnk, error)
	UpdateGroupUser(groupUserParams mysql.UpdateGroupUserLnkParams) error
	DeleteGroupUser(groupUserId int32) error
	GetGroupUser(groupUserId int32) (*gul.GroupsUsersLnk, error)
	ListGroupUsers() ([]*gul.GroupsUsersLnk, error)
}

func (cs *groupUsersService) InsertGroupUser(groupUserParams mysql.CreateGroupUserLnkParams) (*gul.GroupsUsersLnk, error) {
	res, err := gul.InsertGroupUsersLnk(groupUserParams)
	if err != nil {
		return nil, err
	}
	groupUser, err := getGroupUserFromRes(res)
	if err != nil {
		return nil, err
	}

	return groupUser, nil
}

func (cs *groupUsersService) InsertGroupUserByGroupName(userId int32, groupName string) (*gul.GroupsUsersLnk, error) {
	params := mysql.CreateGroupUserLnkByGroupNameParams{
		UserID: userId,
		GName:  sql.NullString{String: groupName, Valid: true},
	}
	res, err := gul.InsertGroupUserLnkByGroupName(params)
	if err != nil {
		return nil, err
	}
	groupUser, err := getGroupUserFromRes(res)
	if err != nil {
		return nil, err
	}
	return groupUser, nil
}

func (cs *groupUsersService) UpdateGroupUser(groupUserParams mysql.UpdateGroupUserLnkParams) error {
	return gul.UpdateGroupUsersLnk(groupUserParams)
}

func (cs *groupUsersService) DeleteGroupUser(groupUserId int32) error {
	return gul.DeleteGroupsUsersLnk(groupUserId)
}

func (cs *groupUsersService) GetGroupUser(groupUserId int32) (*gul.GroupsUsersLnk, error) {
	gu, err := gul.GetGroupUserLnk(groupUserId)
	if err != nil {
		return nil, err
	}
	return convertMysqlGuToGuResponse(gu), nil

}

func (cs *groupUsersService) ListGroupUsers() ([]*gul.GroupsUsersLnk, error) {
	guList, err := gul.ListGroupsUsersLnk()
	if err != nil {
		return nil, err
	}
	return convertMysqlGuListToGu(guList), nil
}

// Utility functions
func getGroupUserFromRes(res sql.Result) (*gul.GroupsUsersLnk, error) {
	groupUserId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	groupUser, err := GroupsUsersService.GetGroupUser(int32(groupUserId))
	if err != nil {
		return nil, err
	}

	return groupUser, nil
}

func convertMysqlGuToGuResponse(gu mysql.GroupsUsersLnk) *gul.GroupsUsersLnk {
	return &gul.GroupsUsersLnk{
		GuID:      gu.GuID,
		GuUserID:  gu.GuUserID.Int32,
		GuGroupID: gu.GuGroupID.Int32,
	}
}

func convertMysqlGuListToGu(mysqlGu []mysql.GroupsUsersLnk) []*gul.GroupsUsersLnk {
	guResponse := []*gul.GroupsUsersLnk{}
	for _, mysqlgu := range mysqlGu {
		guResponse = append(guResponse, convertMysqlGuToGuResponse(mysqlgu))
	}
	return guResponse
}
