package firebase

import (
	"context"
	"log"

	f "firebase.google.com/go"
	"firebase.google.com/go/auth"
	fs "github.com/ravayak/copee_backend/apis/services/firebase"
	"google.golang.org/api/option"
)

var (
	FirebaseService fs.FirebaseService
	FirebaseClient  *auth.Client
)

func FirebaseInit() {
	sa := option.WithCredentialsFile("./datasources/firebase/firebase-service-account-key.json")
	firebaseNewApp, err := f.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalf("Firebase service unabled to start:%v", err)
		panic(err)
	}

	log.Println("Firebase service successfully started")
	FirebaseService.FirebaseApp = firebaseNewApp

	client, err := FirebaseService.FirebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("Unabled to get Firebase client :%v", err)
		panic(err)
	}
	FirebaseClient = client
	log.Println("Firebase client successfully initialized")
}
