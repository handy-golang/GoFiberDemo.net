package midst

import (
	"GoFiberDemo.net/server/global/config"
	"GoFiberDemo.net/server/router/res"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	json := mFiber.DataParser(c)

	ReturnData := make(map[string]any)
	ReturnData["ResParam"] = json
	ReturnData["Method"] = c.Method()
	ReturnData["AppInfo"] = config.AppInfo

	ReturnData["UserAgent"] = c.Get("User-Agent")
	ReturnData["FullPath"] = c.BaseURL() + c.OriginalURL()
	ReturnData["ContentType"] = c.Get("Content-Type")

	// 获取 token
	token := c.Get("Token")
	if len(token) > 0 {
		// 在这里解析 Token,解析正确则返回,否则不返回
		ReturnData["Token"] = token
		return c.JSON(res.OK.WithData(ReturnData))
	} else {
		return c.JSON(res.OK.WithData(ReturnData))
	}
}
