package users_lnk

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ravayak/copee_backend/apis/domain/groups"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGroupsUsersLnkDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	groupsDependencies, err := GetGroupsDependencies()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get group user link
	groupUserLnk, err := GetGroupUserLnk(int32(groupsDependencies.GroupUserLnkId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, groupUserLnk.GuID, int32(groupsDependencies.GroupUserLnkId))

	groupUserLnkByGroupNameParams := mysql.CreateGroupUserLnkByGroupNameParams{
		UserID: int32(groupsDependencies.UserId),
		GName:  sql.NullString{String: "Test Group", Valid: true},
	}
	res, err := InsertGroupUserLnkByGroupName(groupUserLnkByGroupNameParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	groupUserLnkByGroupNameId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupUserLnkByGroupeName, err := GetGroupUserLnk(int32(groupUserLnkByGroupNameId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	assert.Equal(t, groupUserLnkByGroupeName.GuID, int32(groupUserLnkByGroupNameId))

	groupUpdateParams := mysql.UpdateGroupUserLnkParams{
		GuUserID:  sql.NullInt32{Int32: int32(groupsDependencies.UserId), Valid: true},
		GuGroupID: sql.NullInt32{Int32: int32(groupsDependencies.GroupId), Valid: true},
	}

	// Update group user link
	err = UpdateGroupUsersLnk(groupUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated group user link
	groupsUserLnks, err := ListGroupsUsersLnk()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, groupsUserLnks[0].GuGroupID.Int32, int32(groupsDependencies.GroupId))

	// Get group list by user
	groupListByUser, err := groups.ListGroupsByUser(int32(groupsDependencies.UserId))
	assert.Equal(t, groupListByUser[0].GDescription.String, "Test Description")

	err = DeleteGroupDepencies(groupsDependencies)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
