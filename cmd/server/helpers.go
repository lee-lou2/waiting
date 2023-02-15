package server

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

// generateUUID UUID 생성
func generateUUID() (string, error) {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuidObj.String(), nil
}

// generateUrl URL 생성
func generateUrl() string {
	app := os.Getenv("BRAND_NAME")
	host := os.Getenv("PROJECT_HOST")

	uuidObj, _ := generateUUID()
	return fmt.Sprintf("%s/v1/qr/%s/%s/", host, app, uuidObj)
}
