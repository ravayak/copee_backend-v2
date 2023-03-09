package users

import (
	"github.com/ravayak/copee_backend/datasources/mysql"
)

type UserLogin struct {
	User       *mysql.User              `json:"user"`
	LoginTime  string                   `json:"login_time"`
	Privileges *[]mysql.GroupsPrivilege `json:"privileges"`
	Groups     *[]mysql.Group           `json:"groups"`
}
