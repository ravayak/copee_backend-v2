package groups

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/groups"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	GroupsService groupsServiceInterface = &groupsService{}
)

type groupsService struct{}

type groupsServiceInterface interface {
	InsertGroup(groupParams mysql.CreateGroupParams) (sql.Result, error)
	UpdateGroup(groupParams mysql.UpdateGroupParams) error
	DeleteGroup(groupId int32) error
	GetGroup(groupId int32) (mysql.Group, error)
	GetGroupByName(groupName string) (mysql.Group, error)
	ListGroups() ([]mysql.Group, error)
}

func (cs *groupsService) InsertGroup(groupParams mysql.CreateGroupParams) (sql.Result, error) {
	return groups.InsertGroup(groupParams)
}

func (cs *groupsService) UpdateGroup(groupParams mysql.UpdateGroupParams) error {
	return groups.UpdateGroup(groupParams)
}

func (cs *groupsService) DeleteGroup(groupId int32) error {
	return groups.DeleteGroup(groupId)
}

func (cs *groupsService) GetGroup(groupId int32) (mysql.Group, error) {
	return groups.GetGroup(groupId)
}

func (cs *groupsService) GetGroupByName(groupName string) (mysql.Group, error) {
	return groups.GetGroupByName(groupName)
}

func (cs *groupsService) ListGroups() ([]mysql.Group, error) {
	return groups.ListGroups()
}
