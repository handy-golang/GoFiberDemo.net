package public

import (
	"GoFiberDemo.net/server/router/midst"
	"github.com/gofiber/fiber/v2"
)

/*
/api/public
*/

func Router(api fiber.Router) {
	r := api.Group("/public")

	r.Get("", midst.Index("正在访问 /api/public "))
}
