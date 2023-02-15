package server

import "github.com/google/uuid"

// generateUUID UUID 생성
func generateUUID() (string, error) {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return uuidObj.String(), nil
}
