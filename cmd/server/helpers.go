package server

import (
	"fmt"
	"os"
	"strconv"
	"waiting/config/db"
)

// generateUrl URL 생성
func generateUrl() string {
	var codeObj AccessCode

	store := os.Getenv("STORE_ID")
	host := os.Getenv("PROJECT_HOST")

	storeId, _ := strconv.Atoi(store)
	codeObj.StoreId = storeId

	tx, _ := db.GetDatabase()
	tx.Create(&codeObj)
	return fmt.Sprintf("%s/v1/qr/%s/%s/", host, store, codeObj.UUID)
}
