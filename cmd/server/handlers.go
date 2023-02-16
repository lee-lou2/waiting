package server

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"waiting/config/db"
)

// holdingHandler 대기 등록
func holdingHandler(c *fiber.Ctx) error {
	var codeObj AccessCode
	var storeObj Store

	store := c.Params("store")
	storeId, _ := strconv.Atoi(store)
	reqUuid := c.Params("uuid")

	// 0. UUID 가 유효한지 확인
	tx, _ := db.GetDatabase()
	tx.Where(
		"store_id = ? and uuid = ? and is_expired = ?",
		storeId,
		reqUuid,
		false,
	).First(&codeObj)
	if codeObj.UUID == "" {
		return c.Render("error", fiber.Map{})
	}
	// 만료 처리
	codeObj.IsExpired = true
	tx.Save(&codeObj)

	// 1. 다음 QR 코드 생성
	nextUrl := generateUrl(uint(storeId))
	if err := QuePublisher(nextUrl); err != nil {
		return c.Render("error", fiber.Map{})
	}

	// 2. 작성해야하는 데이터 조회
	tx.Preload("Brand").Preload("Forms", func(tx *gorm.DB) *gorm.DB {
		return tx.Where("is_active = ?", true)
	}).Preload("Location").Where("id = ?", storeId).First(&storeObj)
	if storeObj.ID == 0 {
		return c.Render("error", fiber.Map{})
	}
	return c.Render("detail", fiber.Map{
		"store":    storeObj,
		"brand":    storeObj.Brand,
		"forms":    storeObj.Forms,
		"location": storeObj.Location,
	})
}
