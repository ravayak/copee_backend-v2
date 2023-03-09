package context

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	g "github.com/ravayak/copee_backend/apis/domain/groups"
	p "github.com/ravayak/copee_backend/apis/domain/groups/privileges"
	u "github.com/ravayak/copee_backend/apis/domain/users"
	us "github.com/ravayak/copee_backend/apis/services/users"
	errorsdata "github.com/ravayak/copee_backend/app/data/errors"
	f "github.com/ravayak/copee_backend/app/firebase"
	"github.com/ravayak/copee_backend/app/utils/logger"
	stringutils "github.com/ravayak/copee_backend/app/utils/string"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

func GetUserLoginFromContext(c *gin.Context) (u.UserLogin, error) {
	var us u.UserLogin
	userLogin, exists := c.Get("userLogin")
	if exists {
		us = userLogin.(u.UserLogin)
	} else {
		return us, errors.New(errorsdata.UserLoginError)
	}
	return us, nil
}

func GetUserFromContext(c *gin.Context) (*mysql.User, error) {
	token := c.Query("token")
	authToken, err := f.GetAuthToken(token)
	if err != nil {
		return nil, err
	}

	userRecord, err := f.AuthByUid(authToken.UID)
	if err != nil {
		return nil, err
	}

	// Get user in db by mail
	user, err := us.UsersService.GetUserByEmail(userRecord.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUserLoginFromContext(c *gin.Context, user *mysql.User) error {
	groups, err := g.ListGroupsByUser(user.UserID)
	if err != nil {
		logger.Info(fmt.Sprintf("error in getting groups %v for user %d", err, user.UserID))
		return err
	}

	privileges, err := p.GetGroupPrivilegesByUserId(user.UserID)
	if err != nil {
		logger.Info(fmt.Sprintf("error in getting privileges %v for user %d", err, user.UserID))
		return err
	}

	userLogin := u.UserLogin{
		User:       user,
		LoginTime:  time.Now().Format("2006-01-02 15:04:05"),
		Privileges: &privileges,
		Groups:     &groups,
	}
	c.Set("userLogin", userLogin)
	return nil
}

func GetParamId(c *gin.Context, paramId string) (int32, error) {
	idParam := c.Param(paramId)
	if idParam == "" {
		return 0, nil
	}
	id, err := stringutils.StrToInt(idParam)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return int32(id), nil
}
