package router

import (
	"os"
	"time"

	"GoFiberDemo.net/server/global/config"
	"GoFiberDemo.net/server/router/midst"
	"GoFiberDemo.net/server/router/private"
	"GoFiberDemo.net/server/router/public"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	fileName := config.Dir.Log + "/HTTP-" + time.Now().Format("06年1月02日15时") + ".log"
	logFile, _ := os.Create(fileName)

	app := fiber.New(fiber.Config{
		ServerHeader: "GoFiberDemo.net",
	})
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}:${port}] ${status} - ${method} ${path} ${latency} \n",
		TimeFormat: "2006-01-02 - 15:04:05",
		Output:     logFile,
	}))

	app.Get("/", midst.Index("欢迎访问 GoFiberDemo.net 服务"))

	// api
	api := app.Group("/api")
	api.Get("", midst.Index("欢迎访问 /api "))

	// public
	public.Router(api)

	// private
	private.Router(api)

	app.Listen(":3000")
}
