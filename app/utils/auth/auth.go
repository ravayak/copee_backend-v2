package privileges

import (
	u "github.com/ravayak/copee_backend/apis/domain/users"
	gd "github.com/ravayak/copee_backend/app/data/global"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

func FindPrivilege(privileges []mysql.GroupsPrivilege, privilege string) bool {
	for _, p := range privileges {
		if p.GpPrivilege.String == privilege {
			return true
		}
	}
	return false
}

func IsUserAdmin(user u.UserLogin) bool {
	for _, group := range *user.Groups {
		if group.GName.String == gd.ADMIN_STATUS {
			return true
		}
	}
	return false
}
