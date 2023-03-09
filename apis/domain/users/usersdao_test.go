package users

import (
	"database/sql"
	"fmt"
	"testing"

	dt "github.com/ravayak/copee_backend/app/data/tests"
	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestUsersDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")

	// Create a user
	resInsertUser, err := InsertUser(dt.CreateUserParams)
	if err != nil {
		fmt.Println(err)
	}

	userId, err := resInsertUser.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get User
	user, err := GetUser(int32(userId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, user.UserID, int32(userId))

	userUpdateParams := mysql.UpdateUserParams{
		UserEmail:           sql.NullString{String: "testupdated@test.test", Valid: true},
		UserLastname:        sql.NullString{String: "Test updated", Valid: true},
		UserFirstname:       sql.NullString{String: "Test updated", Valid: true},
		UserDateCreated:     user.UserDateCreated,
		UserRecruitmentDate: user.UserRecruitmentDate,
		UserID:              int32(userId),
	}

	// Update User
	err = UpdateUser(userUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated User
	users, err := ListUsers()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, users[0].UserEmail.String, "testupdated@test.test")

	// Delete User
	err = DeleteUser(int32(userId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
