package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	u "github.com/ravayak/copee_backend/apis/domain/users"
	gd "github.com/ravayak/copee_backend/app/data/global"
	md "github.com/ravayak/copee_backend/app/data/middlewares"
	gu "github.com/ravayak/copee_backend/app/utils/gin"
	re "github.com/ravayak/copee_backend/app/utils/resterrors"
	"github.com/ravayak/copee_backend/app/utils/url"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := gu.GetUserFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gd.UNAUTHORIZED)
			c.Abort()
			return
		}

		err = gu.SaveUserLoginFromContext(c, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
			return
		}

		c.Next()
	}
}

func EnsureHasPrivileges() gin.HandlerFunc {
	return func(c *gin.Context) {
		userLogin, _ := c.Get(md.USER_LOGIN)
		ul := userLogin.(u.UserLogin)
		resourceName := url.GetResourceName(c.Request.URL.Path)
		privilege := GetPrivilegeByMethod(c.Request.Method, resourceName)
		var err re.RestError

		switch resourceName {
		case md.CLIENTS:
			err = checkPrivilegesByResourceName(c, ul, privilege, "client_id", CheckIfClientIdBelongsToUserClients)
		case md.GROUPS_USERS:
			err = checkPrivilegesByResourceName(c, ul, privilege, "", nil)
		case md.COMPANIES:
			err = checkPrivilegesByResourceName(c, ul, privilege, "company_id", nil)
		default:
			c.Next()
		}
		if err != nil {
			c.JSON(err.Status(), err)
			c.Abort()
		}
		c.Next()
	}
}
