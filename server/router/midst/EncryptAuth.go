package midst

import (
	"errors"

	"GoFiberDemo.net/server/global/config"
	"github.com/gofiber/fiber/v2"
)

func EncryptAuth(c *fiber.Ctx) error {
	EncStr := c.Get("Auth-Encrypt")

	shaStr := config.Encrypt(c.Path())

	if EncStr != shaStr {
		return errors.New("授权验证错误")
	}

	return nil
}
