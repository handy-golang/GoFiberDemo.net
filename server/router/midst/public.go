package midst

import (
	"github.com/gofiber/fiber/v2"
)

func Public(c *fiber.Ctx) error {
	// 添加访问头
	AddHeader(c)

	// 接口授权验证
	// err := EncryptAuth(c)
	// if err != nil {
	// 	return c.JSON(result.ErrApiAuth.WithData(mStr.ToStr(err)))
	// }

	return c.Next()
}

func AddHeader(c *fiber.Ctx) error {
	c.Response().Header.Del("Data-Type")
	c.Set("Data-Type", "GoFiberDemo.net")
	return nil
}
