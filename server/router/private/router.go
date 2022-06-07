package private

import (
	"GoFiberDemo.net/server/router/midst"
	"github.com/gofiber/fiber/v2"
)

/*

/api/private

*/

func Router(api fiber.Router) {
	r := api.Group("/private")

	r.Get("", midst.Index("正在访问 /api/private "))

	r.Get("/ping", midst.Ping)
	r.Post("/ping", midst.Ping)
}
