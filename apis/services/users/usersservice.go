package clientsservice

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/users"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}
type usersServiceInterface interface {
	InsertUser(userParams mysql.CreateUserParams) (sql.Result, error)
	UpdateUser(userParams mysql.UpdateUserParams) error
	DeleteUser(userId int32) error
	GetUser(userId int32) (mysql.User, error)
	ListUsers() ([]mysql.User, error)
	GetUserByEmail(email string) (mysql.User, error)
}

func (us *usersService) InsertUser(userParams mysql.CreateUserParams) (sql.Result, error) {
	return users.InsertUser(userParams)
}

func (us *usersService) UpdateUser(userParams mysql.UpdateUserParams) error {
	return users.UpdateUser(userParams)
}

func (us *usersService) DeleteUser(userId int32) error {
	return users.DeleteUser(userId)
}

func (us *usersService) GetUser(userId int32) (mysql.User, error) {
	return users.GetUser(userId)
}

func (us *usersService) GetUserByEmail(email string) (mysql.User, error) {
	return users.GetUserByEmail(email)
}

func (us *usersService) ListUsers() ([]mysql.User, error) {
	return users.ListUsers()
}
