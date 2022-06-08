package router

import (
	"net/http"
	"os"
	"time"

	"GoFiberDemo.net/server/global"
	"GoFiberDemo.net/server/global/config"
	"GoFiberDemo.net/server/router/midst"
	"GoFiberDemo.net/server/router/private"
	"GoFiberDemo.net/server/router/public"
	"GoFiberDemo.net/server/tmpl"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	fileName := config.Dir.Log + "/HTTP-" + time.Now().Format("06年1月02日15时") + ".log"
	logFile, _ := os.Create(fileName)

	app := fiber.New(fiber.Config{
		ServerHeader: "GoFiberDemo.net",
	})

	// 日志中间件
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}:${port}] ${status} - ${method} ${path} ${latency} \n",
		TimeFormat: "2006-01-02 - 15:04:05",
		Output:     logFile,
	}))

	// 静态文件服务器
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(tmpl.Static),
	}))

	// api
	api := app.Group("/api")
	api.Get("/ping", midst.Ping)
	api.Post("/ping", midst.Ping)

	// public
	public.Router(api)

	// private
	private.Router(api)

	listenHost := mStr.Join(":", config.AppEnv.Port)
	global.Log.Println(mStr.Join(`启动服务: http://127.0.0.1`, listenHost))
	app.Listen(listenHost)
}
