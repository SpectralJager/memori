package main

import (
	"context"

	"github.com/SpectralJager/memori/view"
	"github.com/labstack/echo/v4"
)

var quotes = []string{
	"Laudantium mollitia deserunt molestiae aperiam ut cum.",
	"Voluptas quia dolor labore.",
	"Earum fuga maiores eos sunt eum laudantium saepe qui id. Velit ducimus architecto maxime enim voluptatem doloribus dignissimos.",
	"Perspiciatis impedit officiis itaque voluptatem quia excepturi vel.",
}

func main() {
	app := echo.New()

	app.Static("/static", "./public")

	app.GET("/", func(c echo.Context) error {
		return view.IndexView(quotes).Render(context.TODO(), c.Response())
	})

	app.Start(":8080")
}
