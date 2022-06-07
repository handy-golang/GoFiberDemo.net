package midst

import (
	"GoFiberDemo.net/server/router/res"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/gofiber/fiber/v2"
)

func GetPing(c *fiber.Ctx) error {
	return PingAction(c)
}

func PostPing(c *fiber.Ctx) error {
	return PingAction(c)
}

func PingAction(c *fiber.Ctx) error {
	// 在这里可以解析参数
	var data struct {
		FullPath string
	}
	json := mFiber.DataParser(c, &data)

	ReturnData := make(map[string]any)
	ReturnData["ResParam"] = json
	ReturnData["Method"] = c.Method()

	return c.JSON(res.OK.WithData(ReturnData))
}
