package private

import (
	"GoFiberDemo.net/server/router/midst"
	"GoFiberDemo.net/server/router/result"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Type", "GoFiberDemo.net/api/private")

	// 授权验证
	err := midst.EncryptAuth(c)
	if err != nil {
		return c.JSON(result.ErrApiAuth.WithData(mStr.ToStr(err)))
	}

	return c.Next()
}
