package midst

import (
	"errors"
	"time"

	"GoFiberDemo.net/server/global/config"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/gofiber/fiber/v2"
)

func TokenAuth(c *fiber.Ctx) (UserID string, err error) {
	Token := c.Get("Token")
	if len(Token) < 1 {
		err = errors.New("缺少Token")
		return
	}

	Claims, AuthOk := mEncrypt.ParseToken(Token, config.SecretKey)
	if !AuthOk {
		err = errors.New("Token验证失败")
		return
	}

	UserID = Claims.Message
	if UserID != "670188307@qq.com" {
		err = errors.New("Token解析失败")
		return
	}

	ExpiresAt := Claims.StandardClaims.ExpiresAt
	now := time.Now().Unix()

	if ExpiresAt-now < 0 {
		err = errors.New("Token过期,请重新登录")
		return
	}

	return UserID, nil
}
