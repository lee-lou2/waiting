package server

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Brand 브랜드
type Brand struct {
	gorm.Model
	Name   string `gorm:"not null;type:varchar(50)" json:"name"`
	Stores []Store
}

// Store 매장
type Store struct {
	gorm.Model
	BrandId  uint   `gorm:"foreignKey:ID"`
	Brand    Brand  `gorm:"constraint:OnDelete:CASCADE;" json:"brand"`
	Name     string `gorm:"not null;type:varchar(50)" json:"name"`
	Forms    []StoreForm
	Location StoreLocation
}

// StoreLocation 매장 위치
type StoreLocation struct {
	gorm.Model
	StoreId   uint
	Latitude  string `gorm:"null" json:"latitude"`
	Longitude string `gorm:"null" json:"longitude"`
	IsActive  bool   `gorm:"default:false"`
}

// StoreForm 스토어 작성 양식
type StoreForm struct {
	gorm.Model
	StoreId  uint   `gorm:"foreignKey:ID"`
	Store    Store  `gorm:"constraint:OnDelete:CASCADE;" json:"store"`
	Key      string `gorm:"not null;type:varchar(50)" json:"key"`
	Type     string `gorm:"not null;type:varchar(50)" json:"type"`
	Required bool   `gorm:"default:false" json:"required"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}

// AccessCode 접속 코드
type AccessCode struct {
	UUID      string `gorm:"primarykey;type:uuid"`
	CreatedAt time.Time
	StoreId   uint  `gorm:"foreignKey:ID"`
	Store     Store `gorm:"constraint:OnDelete:CASCADE;" json:"store"`
	IsExpired bool  `gorm:"default:false" json:"is_expired"`
}

// BeforeCreate UUID 생성
func (a *AccessCode) BeforeCreate(tx *gorm.DB) error {
	// UUID 생성
	_uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	a.UUID = _uuid.String()
	return nil
}
