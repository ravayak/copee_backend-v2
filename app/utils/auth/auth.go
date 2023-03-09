package privileges

import (
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
