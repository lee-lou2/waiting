package client

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"time"
)

// qrPageHandler 기본 페이지
func qrPageHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

// streamHandler 스트림 핸들러
func streamHandler(c *fiber.Ctx) error {
	ctx := c.Context()

	ctx.SetContentType("text/event-stream")
	ctx.Response.Header.Set("Cache-Control", "no-cache")
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.Response.Header.Set("Transfer-Encoding", "chunked")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Cache-Control")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for {
			msg := <-QRChan
			fmt.Println("채널이 들어왔습니다 : " + msg)
			fmt.Fprintf(w, "data: Message: %s\n\n", msg)
			w.Flush()
			time.Sleep(time.Second)
		}
	}))
	return nil
}
