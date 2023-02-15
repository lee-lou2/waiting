package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"waiting/config/env"
)

func Run() {
	// 환경 설정
	env.Load()

	// API 실행
	r := gin.Default()
	r.LoadHTMLGlob("./web/*")

	v1 := r.Group("/v1")
	qr := v1.Group("/qr")
	{
		qr.GET("/:app/:uuid/", func(c *gin.Context) {
			app := c.Param("app")
			uuid := c.Param("uuid")

			uuidObj, _ := generateUUID()
			err := messageQuePublisher(uuid)
			fmt.Println(err)
			// 1. 다음 QR 코드 생성
			// 2. 작성해야하는 데이터 조회
			c.HTML(200, "detail.html", gin.H{"app": app, "uuid": uuid})
		})
	}

	r.Run(":8000")
}
