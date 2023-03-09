package firebase

import (
	"context"

	"firebase.google.com/go/auth"
	f "github.com/ravayak/copee_backend/datasources/firebase"
)

func GetAuthToken(token string) (*auth.Token, error) {
	authToken, err := f.FirebaseClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, err
	}

	return authToken, nil
}

func AuthByUid(uid string) (*auth.UserRecord, error) {
	user, err := f.FirebaseClient.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUIDFromToken(token string) (string, error) {
	authToken, err := GetAuthToken(token)
	if err != nil {
		return "", err
	}

	return authToken.UID, nil
}
