package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"waiting/config/db"
)

// holdingHandler 대기 등록
func holdingHandler(c *fiber.Ctx) error {
	var codeObj AccessCode
	var storeObj Store
	var storeLocationObj StoreLocation
	//var storeFormObj []StoreForm

	store := c.Params("store")
	reqUuid := c.Params("uuid")

	// 0. UUID 가 유효한지 확인
	tx, _ := db.GetDatabase()
	tx.Where(&AccessCode{UUID: reqUuid, IsExpired: false}).First(&codeObj)
	//if codeObj.UUID == "" {
	//	return c.Render("error", fiber.Map{})
	//}
	//// 만료 처리
	//codeObj.IsExpired = true
	//tx.Save(&codeObj)

	// 1. 다음 QR 코드 생성
	nextUrl := generateUrl()
	fmt.Println(nextUrl)
	if err := QuePublisher(nextUrl); err != nil {
		return c.Render("error", fiber.Map{})
	}

	// 2. 작성해야하는 데이터 조회
	// 스토어 조회
	storeId, _ := strconv.Atoi(store)
	tx.Preload("Forms").Where("id = ?", storeId).First(&storeObj)
	if storeObj.ID == 0 {
		return c.Render("error", fiber.Map{})
	}
	for _, form := range storeObj.Forms {
		fmt.Println(form.Key)
	}

	// 지역 정보 조회
	tx.Where(&StoreLocation{Store: storeObj}).Find(&storeLocationObj)

	return c.Render("detail", fiber.Map{
		"store_name": store,
	})
}
