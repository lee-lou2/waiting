package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"waiting/config/db"
)

func Run() {
	tx, _ := db.GetDatabase()
	tx.AutoMigrate(&Brand{}, &Store{}, &StoreLocation{}, &StoreForm{})

	engine := html.New("./web", ".html")
	app := fiber.New(
		fiber.Config{Views: engine},
	)
	app.Use(logger.New())

	v1 := app.Group("/v1")
	qr := v1.Group("/qr")
	{
		qr.Get("/:app/:uuid", holdingHandler)
	}

	app.Listen(":3001")
}
