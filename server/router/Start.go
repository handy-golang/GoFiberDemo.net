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
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func Start() {
	// 加载日志文件
	fileName := config.Dir.Log + "/HTTP-" + time.Now().Format("06年1月02日15时") + ".log"
	logFile, _ := os.Create(fileName)
	/*
		加载模板
		https://www.gouguoyin.cn/posts/10103.html
	*/
	engine := html.NewFileSystem(http.FS(tmpl.Static), ".html")

	// 创建服务
	app := fiber.New(fiber.Config{
		ServerHeader: "GoFiberDemo.net",
		Views:        engine,
	})

	// 跨域
	app.Use(cors.New())

	// 限流
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 20 * time.Second,
	}))

	// 日志中间件
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}:${port}] ${status} - ${method} ${latency} ${path} \n",
		TimeFormat: "2006-01-02 - 15:04:05",
		Output:     logFile,
	}), midst.Public, compress.New(), favicon.New())

	// 模板渲染样例
	app.Get("/", Index)

	// api
	api := app.Group("/api")
	api.Get("/ping", midst.Ping)
	api.Post("/ping", midst.Ping)

	// public
	public.Router(api)

	// private
	private.Router(api)

	// 静态文件服务器
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(tmpl.Static),
	}))

	// 404
	app.Use(NotFund)

	listenHost := mStr.Join(":", config.AppEnv.Port)
	global.Log.Println(mStr.Join(`启动服务: http://127.0.0.1`, listenHost))
	app.Listen(listenHost)
}
