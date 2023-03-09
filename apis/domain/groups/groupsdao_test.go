package groups

import (
	"database/sql"
	"fmt"
	"testing"

	dt "github.com/ravayak/copee_backend/app/data/tests"
	mysqlutils "github.com/ravayak/copee_backend/app/utils/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGroupsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")
	mysqlutils.DeleteDatasFromDB()

	// Create group
	CreateGroupParams := mysql.CreateGroupParams{
		GName:        sql.NullString{String: "Test", Valid: true},
		GDescription: sql.NullString{String: "Test", Valid: true},
	}

	resInsertGroup, err := InsertGroup(CreateGroupParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	groupId, err := resInsertGroup.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get group
	group, err := GetGroup(int32(groupId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, group.GID, int32(groupId))

	// Get group by name
	groupByName, err := GetGroupByName("Test")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, groupByName.GName.String, "Test")

	groupUpdateParams := dt.GetGroupUpdateParams(int32(groupId))

	// Update group
	err = UpdateGroup(groupUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated group
	groups, err := ListGroups()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, groups[0].GName.String, "Test updated")

	// Delete Group
	err = DeleteGroup(int32(groupId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
