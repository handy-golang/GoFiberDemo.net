package midst

import (
	"GoFiberDemo.net/server/router/res"
	"github.com/gofiber/fiber/v2"
)

func GetPing(c *fiber.Ctx) {
	json := make(map[string]string)

	PingAction(c, json)
}

func PostPing(c *fiber.Ctx) {
	json := make(map[string]string)

	PingAction(c, json)
}

func PingAction(c *fiber.Ctx, json map[string]string) {
	ReturnData := make(map[string]any)
	ReturnData["ResParam"] = json

	c.JSON(res.OK.WithData(ReturnData))
}
