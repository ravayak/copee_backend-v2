package installationsdependencies

import (
	"github.com/ravayak/copee_backend/apis/domain/clients"
	"github.com/ravayak/copee_backend/apis/domain/users"
	dt "github.com/ravayak/copee_backend/app/data/tests"
)

func CreateInstallationDepencencies() (int64, int64, int64, error) {
	// Create client to fullfil Foreign Key constraint
	resInsertClient, err := clients.InsertClient(dt.CreateClientParams)
	if err != nil {
		return -1, -1, -1, err
	}

	// Create user to fullfil Foreign Key constraint
	resCreateUser, err := users.InsertUser(dt.CreateUserParams)
	if err != nil {
		return -1, -1, -1, err
	}

	// Create shared user to fullfil Foreign Key constraint
	resCreateSharedUser, err := users.InsertUser(dt.CreateSharedUserParams)
	if err != nil {
		return -1, -1, -1, err
	}

	clientId, err := resInsertClient.LastInsertId()
	if err != nil {
		return -1, -1, -1, err
	}

	userId, err := resCreateUser.LastInsertId()
	if err != nil {
		return -1, -1, -1, err
	}

	sharedUserId, err := resCreateSharedUser.LastInsertId()
	if err != nil {
		return -1, -1, -1, err
	}

	return clientId, userId, sharedUserId, nil

}

func DeleteInstallationDependencies(userId, sharedUserId, clientId int32) error {
	// Delete User
	err := users.DeleteUser(userId)
	if err != nil {
		return err
	}

	// Delete Shared User
	err = users.DeleteUser(sharedUserId)
	if err != nil {
		return err
	}
	// Delete Client
	err = clients.DeleteClient(clientId)
	if err != nil {
		return err
	}
	return nil
}
