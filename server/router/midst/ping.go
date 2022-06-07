package midst

import (
	"GoFiberDemo.net/server/router/res"
	"github.com/gofiber/fiber/v2"
)

func GetPing(c *fiber.Ctx) error {
	json := make(map[string]string)

	return PingAction(c, json)
}

func PostPing(c *fiber.Ctx) error {
	json := make(map[string]string)
	c.QueryString()

	return PingAction(c, json)
}

func PingAction(c *fiber.Ctx, json map[string]string) error {
	ReturnData := make(map[string]any)
	ReturnData["ResParam"] = json

	return c.JSON(res.OK.WithData(ReturnData))
}
