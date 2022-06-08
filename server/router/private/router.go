package private

import (
	"GoFiberDemo.net/server/router/midst"
	"github.com/gofiber/fiber/v2"
)

/*

/api/private

*/

func Router(api fiber.Router) {
	r := api.Group("/private", MiddleWare)

	r.Get("/ping", midst.Ping)
	r.Post("/ping", midst.Ping)
}
