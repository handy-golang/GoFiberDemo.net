package public

import (
	"time"

	"GoFiberDemo.net/server/global/config"
	"GoFiberDemo.net/server/router/result"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mRes/mFiber"
	"github.com/gofiber/fiber/v2"
)

func GenerateToken(c *fiber.Ctx) error {
	var json struct {
		Email    string `bson:"Email"`    // 670188307@qq.com
		Password string `bson:"Password"` // 密码
	}

	mFiber.DataParser(c, &json)

	if json.Email != "670188307@qq.com" {
		return c.JSON(result.ErrUserNot.WithData("当前账号不存在"))
	}

	if json.Password != "asdasd55555" {
		return c.JSON(result.ErrPassword.WithData("密码不正确"))
	}

	var loginSucceedData struct {
		Email  string `bson:"Email"`
		UserID string `bson:"UserID"`
		Token  string `bson:"Token"`
	}
	loginSucceedData.Email = json.Email
	loginSucceedData.UserID = json.Password
	// 颁发 Token
	loginSucceedData.Token = mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: config.SecretKey,
		ExpiresAt: time.Now().Add(72 * time.Hour),
		Message:   json.Email,
		Issuer:    "mo7.cc",
		Subject:   "UserToken",
	}).Generate()

	return c.JSON(result.ErrPassword.WithData(loginSucceedData))
}
