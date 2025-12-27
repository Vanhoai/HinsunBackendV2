package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	googleFirebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func NewFirebase(ctx context.Context) (*googleFirebase.App, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory:", err)
	}

	servicePath := fmt.Sprintf("%s/%s", rootPath, "firebase-adminsdk.json")
	opt := option.WithCredentialsFile(servicePath)
	app, err := googleFirebase.NewApp(context.Background(), &googleFirebase.Config{}, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}
