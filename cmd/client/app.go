package client

import (
	"github.com/gin-gonic/gin"
	"os"
	"waiting/config/env"
)

func Run() {
	// 환경 설정
	env.Load()

	// 리시버 실행
	go messageQueConsumer()

	// API 실행
	r := gin.Default()

	user := os.Getenv("USER_ID")
	password := os.Getenv("USER_PASSWORD")
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{user: password}))
	authorized.GET("/stream", HeadersMiddleware(), stream.serveHTTP(), streamHandler)

	r.StaticFile("/", "./web/index.html")
	r.Run(":8001")
}
