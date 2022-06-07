package router

import (
	"GoFiberDemo.net/server/router/midst"
	"GoFiberDemo.net/server/router/private"
	"GoFiberDemo.net/server/router/public"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Get("/", midst.Index("欢迎访问 GoFiberDemo.net 服务"))

	// api
	api := app.Group("/api")
	api.Get("", midst.Index("正在访问 /api "))

	// public
	public.Router(api)

	// private
	private.Router(api)

	app.Listen(":3000")
}
