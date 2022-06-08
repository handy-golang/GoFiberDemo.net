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

	r.Get("/ping", midst.Ping)
	r.Post("/ping", midst.Ping)
}
