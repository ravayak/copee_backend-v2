package groups

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/groups/privileges"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	GroupsPrivilegesService groupsPrivilegesServiceInterface = &groupsPrivilegesService{}
)

type groupsPrivilegesService struct{}

type groupsPrivilegesServiceInterface interface {
	InsertGroupPrivilege(privilege string) (sql.Result, error) // Ou est le parametre createGroupPrivilege ?
	DeleteGroupPrivilege(groupPrivilegeId int32) error
	GetGroupPrivilege(groupPrivilegeId int32) (mysql.GroupsPrivilege, error)
	ListGroupPrivileges() ([]mysql.GroupsPrivilege, error)
}

func (cs *groupsPrivilegesService) InsertGroupPrivilege(privilege string) (sql.Result, error) {
	return privileges.InsertGroupPrivilege(sql.NullString{String: privilege, Valid: true})
}

func (cs *groupsPrivilegesService) UpdateGroupPrivilege(groupPrivilegeParams mysql.UpdateGroupPrivilegeParams) error {
	return privileges.UpdateGroupPrivilege(groupPrivilegeParams)
}

func (cs *groupsPrivilegesService) DeleteGroupPrivilege(groupPrivilegeId int32) error {
	return privileges.DeleteGroupPrivilege(groupPrivilegeId)
}

func (cs *groupsPrivilegesService) GetGroupPrivilege(groupPrivilegeId int32) (mysql.GroupsPrivilege, error) {
	return privileges.GetGroupPrivilege(groupPrivilegeId)
}

func (cs *groupsPrivilegesService) ListGroupPrivileges() ([]mysql.GroupsPrivilege, error) {
	return privileges.ListGroupsPrivileges()
}
