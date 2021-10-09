package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-resty/resty/v2"
)

type AppCtx struct {
	app *fiber.App
	client *resty.Client
}

func main() {
	appCtx := AppCtx{
		app: fiber.New(),
		client: resty.New(),
	}
	pokemonRoutes(appCtx)
	appCtx.app.Listen(":3000")
}