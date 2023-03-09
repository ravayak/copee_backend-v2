package middlewares

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	u "github.com/ravayak/copee_backend/apis/domain/users"
	cs "github.com/ravayak/copee_backend/apis/services/clients"
	md "github.com/ravayak/copee_backend/app/data/middlewares"
	authutils "github.com/ravayak/copee_backend/app/utils/auth"
	gu "github.com/ravayak/copee_backend/app/utils/gin"
	re "github.com/ravayak/copee_backend/app/utils/resterrors"
)

// UTILITY FUNCTIONS
func checkPrivilegesByResourceName(c *gin.Context, ul u.UserLogin, privilege, paramId string,
	check func(userLogin u.UserLogin, paramClientId int32) (bool, error)) re.RestError {
	privilegeParams := &md.PrivilegeParams{
		C:             c,
		RequestMethod: c.Request.Method,
		Privilege:     privilege,
		HasPrivilege:  false,
		Ul:            ul,
		ParamId:       paramId,
		Check:         check,
	}
	return CheckPrivileges(c, ul, privilegeParams)
}

func CheckIfClientIdBelongsToUserClients(userLogin u.UserLogin, paramClientId int32) (bool, error) {
	clientsIds, err := cs.ClientsService.ListClientsIdByUser(userLogin.User.UserID)
	if err != nil {
		return false, err
	}

	for _, clientId := range clientsIds {
		if paramClientId == clientId {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("client id %d not valid for user %d", paramClientId, userLogin.User.UserID))
}

func requestMethodMatchMethod(hasPrivilegeParams *md.PrivilegeParams, method string) bool {
	return hasPrivilegeParams.C.Request.Method == method
}

func CheckPrivileges(c *gin.Context, ul u.UserLogin, privilegeParams *md.PrivilegeParams) re.RestError {
	var err error

	privilegeParams, err = EnsurehasPrivilege(privilegeParams)
	if requestMethodMatchMethod(privilegeParams, md.POST) {
		privilegeParams.RequestMethod = md.POST
		privilegeParams, err = EnsurehasPrivilege(privilegeParams)
	}
	if requestMethodMatchMethod(privilegeParams, md.PUT) {
		privilegeParams.RequestMethod = md.PUT
		privilegeParams, err = EnsurehasPrivilege(privilegeParams)
	}
	if requestMethodMatchMethod(privilegeParams, md.DELETE) {
		privilegeParams.RequestMethod = md.DELETE
		privilegeParams, err = EnsurehasPrivilege(privilegeParams)
	}
	if err != nil {
		return re.NewBadRequestError(fmt.Sprintf("error with %v method in route %v for user %d", privilegeParams.C.Request.Method, privilegeParams.C.Request.URL.Path, privilegeParams.Ul.User.UserID), err)
	}
	if !privilegeParams.HasPrivilege {
		return re.NewUnauthorizedError("error checking privileges", errors.New("user does not have privileges"))
	}

	return nil
}

func EnsurehasPrivilege(hasPrivilegeParams *md.PrivilegeParams) (*md.PrivilegeParams, error) {
	if authutils.FindPrivilege(*hasPrivilegeParams.Ul.Privileges, hasPrivilegeParams.Privilege) {
		id, err := gu.GetParamId(hasPrivilegeParams.C, hasPrivilegeParams.ParamId)
		if err != nil {
			return hasPrivilegeParams, err
		}
		if id > 0 {
			if hasPrivilegeParams.Check != nil {
				hasPrivilegeParams.HasPrivilege, err = hasPrivilegeParams.Check(hasPrivilegeParams.Ul, id)
				if err != nil {
					return hasPrivilegeParams, err
				}
			} else {
				hasPrivilegeParams.HasPrivilege = true
			}
			hasPrivilegeParams.C.Set(hasPrivilegeParams.ParamId, id)
		} else {
			hasPrivilegeParams.HasPrivilege = true
		}
	}

	return hasPrivilegeParams, nil
}

func GetPrivilegeByMethod(method, resourceName string) string {
	if strings.Contains(resourceName, md.GROUPS) {
		resourceNameSplit := strings.Split(resourceName, "_")
		resourceName = resourceNameSplit[0]
	}
	privilege := resourceName + "_"

	if method == md.GET {
		privilege += md.READ
	}
	if method == md.POST {
		privilege += md.CREATE
	}
	if method == md.PUT || method == md.PATCH {
		privilege += md.UPDATE
	}
	if method == md.DELETE {
		privilege += md.DELETE
	}
	return strings.ToUpper(privilege)
}
