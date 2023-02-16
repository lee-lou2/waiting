package server

import (
	"fmt"
	"os"
	"waiting/config/db"
)

// generateUrl URL 생성
func generateUrl(storeId uint) string {
	host := os.Getenv("PROJECT_HOST")

	tx, _ := db.GetDatabase()
	codeObj := AccessCode{StoreId: storeId}
	tx.Create(&codeObj)
	return fmt.Sprintf("%s/v1/qr/%d/%s/", host, storeId, codeObj.UUID)
}
