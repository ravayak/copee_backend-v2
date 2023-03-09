package privileges_lnk

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/groups"
	"github.com/ravayak/copee_backend/apis/domain/groups/privileges"
	dt "github.com/ravayak/copee_backend/app/data/tests"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGroupsPrivilegesLnkDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")

	// Create Privilege to fulfill foreign key constraint
	privilege, err := privileges.InsertGroupPrivilege(sql.NullString{String: "test", Valid: true})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	privId, err := privilege.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Create Group to fulfill foreign key constraint
	group, err := groups.InsertGroup(dt.CreateGroupParams)
	groupId, err := group.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	privilegeLnkParams := mysql.CreateGroupPrivilegeLnkParams{
		GplPrivilegeID: sql.NullInt32{Int32: int32(privId), Valid: true},
		GplGroupID:     sql.NullInt32{Int32: int32(groupId), Valid: true},
	}

	insertGroupPrivilegeLnk, err := InsertGroupPrivilegeLnk(privilegeLnkParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	insertGroupPrivilegeLnkId, err := insertGroupPrivilegeLnk.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get group user link
	groupPrivilegeLnk, err := GetGroupPrivilegeLnk(int32(insertGroupPrivilegeLnkId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, groupPrivilegeLnk.GplID, int32(insertGroupPrivilegeLnkId))

	updateGroupUserLnkParams := mysql.UpdateGroupPrivilegeLnkParams{
		GplGroupID:     sql.NullInt32{Int32: int32(groupId), Valid: true},
		GplPrivilegeID: sql.NullInt32{Int32: int32(privId), Valid: true},
		GplID:          int32(insertGroupPrivilegeLnkId),
	}

	// Update group user link
	err = UpdateGroupPrivilegeLnk(updateGroupUserLnkParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated group user link
	groupsUserLnks, err := ListGroupsPrivilegesLnk()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, groupsUserLnks[0].GplGroupID.Int32, int32(groupId))
	assert.Equal(t, groupsUserLnks[0].GplPrivilegeID.Int32, int32(privId))

	// Delete Group
	err = groups.DeleteGroup(int32(groupId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Delete Privilege
	err = privileges.DeleteGroupPrivilege(int32(privId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Delete Group User Link
	err = DeleteGroupPrivilegeLnk(int32(insertGroupPrivilegeLnkId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
