package users

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListUsers() ([]mysql.User, error) {
	ctx := context.Background()
	users, err := mysql.QueriesDb.ListUsers(ctx)
	return users, err
}

func GetUser(userId int32) (mysql.User, error) {
	ctx := context.Background()
	user, err := mysql.QueriesDb.GetUser(ctx, userId)
	return user, err
}

func InsertUser(userParams mysql.CreateUserParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateUser(ctx, userParams)
	return res, err
}

func UpdateUser(userParams mysql.UpdateUserParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateUser(ctx, userParams)
	return err
}

func DeleteUser(userId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteUser(ctx, userId)
	return err
}

func GetUserByEmail(email string) (mysql.User, error) {
	ctx := context.Background()
	user, err := mysql.QueriesDb.GetUserByEmail(ctx, sql.NullString{String: email, Valid: true})
	return user, err
}
