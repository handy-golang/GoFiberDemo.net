package router

import (
	"GoFiberDemo.net/server/global/config"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	// Render index
	return c.Render("index/index.tmpl", fiber.Map{
		"Title":   "GoFiberDemo.net",
		"AppInfo": config.AppInfo,
		"BtnList": []struct {
			Name string
			ID   int
		}{
			{
				Name: "按钮1",
				ID:   1,
			}, {
				Name: "按钮2",
				ID:   2,
			}, {
				Name: "按钮3",
				ID:   3,
			},
		},
	})
}

func NotFund(c *fiber.Ctx) error {
	return c.Status(200).Render("404/404.tmpl", fiber.Map{
		"title": "404 not found",
	})
}

// 404 返回 index.html
// app.Use(func(c *fiber.Ctx) error {
// 	c.Set("Content-Type", "text/html")
// 	return c.Send(tmpl.IndexHtml)
// })
