package client

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

// QRChan QR 기본 채널
var QRChan chan string

// Run 클라이언트 실행
func Run() {
	QRChan = make(chan string)
	go QueConsumer(&QRChan)

	engine := html.New("./web", ".html")
	app := fiber.New(
		fiber.Config{Views: engine},
	)
	app.Use(logger.New())
	app.Get("/", qrPageHandler)
	app.Get("/stream", streamHandler)

	app.Listen(":3000")
}
