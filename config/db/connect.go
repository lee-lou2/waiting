package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var sqliteClient *gorm.DB

// GetDatabase 데이터베이스 연결
func GetDatabase() (*gorm.DB, error) {
	var err error
	if sqliteClient == nil {
		// 데이터베이스 생성
		sqliteClient, err = gorm.Open(
			sqlite.Open(os.Getenv("DATABASE_HOST_SQLITE")), &gorm.Config{},
		)
	}
	if err != nil {
		return nil, err
	}
	return sqliteClient, err
}
