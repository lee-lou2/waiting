package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"waiting/config/db"
)

// holdingHandler 대기 등록
func holdingHandler(c *fiber.Ctx) error {
	var brandObj Brand
	var storeObj Store
	var storeLocationObj StoreLocation
	var storeFormObj []StoreForm

	app := c.Params("app")
	reqUuid := c.Params("uuid")
	fmt.Println(reqUuid)

	// 1. 다음 QR 코드 생성
	nextUrl := generateUrl()
	if err := QuePublisher(nextUrl); err != nil {
		return c.Render("error", fiber.Map{})
	}

	// 2. 작성해야하는 데이터 조회
	tx, _ := db.GetDatabase()
	// 브랜드 조회
	tx.Where(&Brand{Name: app}).First(&brandObj)
	if brandObj.ID == 0 {
		return c.Render("error", fiber.Map{})
	}
	// 스토어 조회
	tx.Where(&Store{Brand: brandObj}).First(&storeObj)
	if storeObj.ID == 0 {
		return c.Render("error", fiber.Map{})
	}

	// 폼 리스트 조회
	tx.Where(&StoreForm{Store: storeObj}).Find(&storeFormObj)
	for _, form := range storeFormObj {
		fmt.Println(form.Key)
		fmt.Println(form.Type)
	}

	// 지역 정보 조회
	tx.Where(&StoreLocation{Store: storeObj}).First(&storeLocationObj)
	fmt.Println(storeLocationObj.ID)

	return c.Render("detail", fiber.Map{
		"brand_name": brandObj.Name,
		"store_name": storeObj.Name,
	})
}
