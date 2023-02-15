package server

import "gorm.io/gorm"

// Brand 브랜드
type Brand struct {
	gorm.Model
	Name string `gorm:"not null;type:varchar(50)" json:"name"`
}

// Store 매장
type Store struct {
	gorm.Model
	BrandId int
	Brand   Brand  `gorm:"constraint:OnDelete:CASCADE;" json:"brand"`
	Name    string `gorm:"not null;type:varchar(50)" json:"name"`
}

// StoreLocation 매장 위치
type StoreLocation struct {
	gorm.Model
	StoreId   int
	Store     Store  `gorm:"constraint:OnDelete:CASCADE;" json:"store"`
	Latitude  string `gorm:"null" json:"latitude"`
	Longitude string `gorm:"null" json:"longitude"`
	IsActive  bool   `gorm:"default:false"`
}

// StoreForm 스토어 작성 양식
type StoreForm struct {
	gorm.Model
	StoreId  int
	Store    Store  `gorm:"constraint:OnDelete:CASCADE;" json:"store"`
	Key      string `gorm:"not null;type:varchar(50)" json:"key"`
	Type     string `gorm:"not null;type:varchar(50)" json:"type"`
	IsActive bool   `gorm:"default:true"`
}
