package midst

import (
	"github.com/gofiber/fiber/v2"
)

func Public(c *fiber.Ctx) error {
	// 添加访问头
	AddHeader(c)

	return c.Next()
}

func AddHeader(c *fiber.Ctx) error {
	c.Set("Data-Path", "GoFiberDemo.net")
	return nil
}
