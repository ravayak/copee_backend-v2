package privileges

import (
	"database/sql"
	"fmt"
	"testing"

	gpl "github.com/ravayak/copee_backend/apis/domain/groups/privileges/lnk"
	gul "github.com/ravayak/copee_backend/apis/domain/groups/users/lnk"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGroupsPrivilegesDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")

	// Create Privilege
	resInsertPrivilege, err := InsertGroupPrivilege(sql.NullString{String: "Privilege", Valid: true})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	privId, err := resInsertPrivilege.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Privilege
	privilege, err := GetGroupPrivilege(int32(privId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, privilege.GpPrivilegeID, int32(privId))

	groupUpdateParams := mysql.UpdateGroupPrivilegeParams{
		GpPrivilegeID: int32(privId),
		GpPrivilege:   sql.NullString{String: "Privilege updated", Valid: true},
	}

	// Update Privilege
	err = UpdateGroupPrivilege(groupUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated Privilege
	privileges, err := ListGroupsPrivileges()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, privileges[0].GpPrivilege.String, "Privilege updated")

	// Create user / group / group user lnk
	groupsDependencies, err := gul.GetGroupsDependencies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Create Group Privilege Lnk
	privilegeLnkParams := mysql.CreateGroupPrivilegeLnkParams{
		GplPrivilegeID: sql.NullInt32{Int32: int32(privId), Valid: true},
		GplGroupID:     sql.NullInt32{Int32: int32(groupsDependencies.GroupId), Valid: true},
	}

	_, err = gpl.InsertGroupPrivilegeLnk(privilegeLnkParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	privilegesByUserId, err := GetGroupPrivilegesByUserId(int32(groupsDependencies.UserId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, privilegesByUserId[0].GpPrivilege.String, "Privilege updated")

	err = DeleteGroupPrivilege(privilege.GpPrivilegeID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = gul.DeleteGroupDepencies(groupsDependencies)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
