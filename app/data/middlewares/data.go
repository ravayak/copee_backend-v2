package middlewares_utils_data

import (
	"github.com/gin-gonic/gin"
	u "github.com/ravayak/copee_backend/apis/domain/users"
)

type PrivilegeParams struct {
	C             *gin.Context
	RequestMethod string
	Privilege     string
	HasPrivilege  bool
	Ul            u.UserLogin
	ParamId       string
	Check         func(u.UserLogin, int32) (bool, error)
}

const POST = "POST"
const READ = "READ"
const PUT = "PUT"
const PATCH = "PATCH"
const GET = "GET"
const UPDATE = "UPDATE"
const CREATE = "CREATE"
const DELETE = "DELETE"
const CLIENTS = "clients"
const USERS = "users"
const GROUPS = "groups"
const GROUPS_USERS = "groups_users"
const CLIENT_GET = "CLIENT_GET"
const CLIENT_LIST = "CLIENT_LIST"
const CLIENT_UPDATE = "CLIENT_UPDATE"
const CLIENT_DELETE = "CLIENT_DELETE"
const USER_LOGIN = "userLogin"
