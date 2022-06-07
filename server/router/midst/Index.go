package midst

import (
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
)

var TempStr = `
<a href="//mo7.cc"> ${Content} </a>
`

func Index(Content string) func(*fiber.Ctx) error {
	dev := map[string]string{
		"Content": Content,
	}

	Str := mStr.Temp(TempStr, dev)

	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.Send([]byte(Str))
	}
}
