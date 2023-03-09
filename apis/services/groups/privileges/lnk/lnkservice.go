package groups

import (
	"database/sql"

	lnk "github.com/ravayak/copee_backend/apis/domain/groups/privileges/lnk"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	GroupsPrivilegesLnkService groupsPrivilegesLnkServiceInterface = &groupsPrivilegesLnkService{}
)

type groupsPrivilegesLnkService struct{}
type groupsPrivilegesLnkServiceInterface interface {
	InsertGroupPrivilegeLnk(groupPrivilegeLnkParams mysql.CreateGroupPrivilegeLnkParams) (sql.Result, error)
	UpdateGroupPrivilegeLnk(groupPrivilegeLnkParams mysql.UpdateGroupPrivilegeLnkParams) error
	DeleteGroupPrivilegeLnk(groupPrivilegeLnkId int32) error
	GetGroupPrivilegeLnk(groupPrivilegeLnkId int32) (mysql.GroupsPrivilegesLnk, error)
	ListGroupPrivilegeLnks() ([]mysql.GroupsPrivilegesLnk, error)
}

func (cs *groupsPrivilegesLnkService) InsertGroupPrivilegeLnk(groupPrivilegeLnkParams mysql.CreateGroupPrivilegeLnkParams) (sql.Result, error) {
	return lnk.InsertGroupPrivilegeLnk(groupPrivilegeLnkParams)
}

func (cs *groupsPrivilegesLnkService) UpdateGroupPrivilegeLnk(groupPrivilegeLnkParams mysql.UpdateGroupPrivilegeLnkParams) error {
	return lnk.UpdateGroupPrivilegeLnk(groupPrivilegeLnkParams)
}

func (cs *groupsPrivilegesLnkService) DeleteGroupPrivilegeLnk(groupPrivilegeLnkId int32) error {
	return lnk.DeleteGroupPrivilegeLnk(groupPrivilegeLnkId)
}

func (cs *groupsPrivilegesLnkService) GetGroupPrivilegeLnk(groupPrivilegeLnkId int32) (mysql.GroupsPrivilegesLnk, error) {
	return lnk.GetGroupPrivilegeLnk(groupPrivilegeLnkId)
}

func (cs *groupsPrivilegesLnkService) ListGroupPrivilegeLnks() ([]mysql.GroupsPrivilegesLnk, error) {
	return lnk.ListGroupsPrivilegesLnk()
}
